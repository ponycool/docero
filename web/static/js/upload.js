// web/static/js/upload.js

const uploadForm = document.getElementById('uploadForm');
const fileInput = document.getElementById('fileInput');
const submitButton = document.getElementById('submitButton');
const responseMessage = document.getElementById('responseMessage');
const downloadLink = document.getElementById('downloadLink');
const loadingSpinner = document.getElementById('loadingSpinner');

// Function to show loading state
function showLoading() {
    submitButton.disabled = true;
    loadingSpinner.style.display = 'block';
    responseMessage.style.display = 'block';
    responseMessage.className = 'message info';
    responseMessage.textContent = 'Uploading and converting, please wait...';
    downloadLink.style.display = 'none';
}

// Function to hide loading state
function hideLoading() {
    submitButton.disabled = false;
    loadingSpinner.style.display = 'none';
}

uploadForm.addEventListener('submit', async function(event) {
    event.preventDefault();

    responseMessage.style.display = 'none';
    downloadLink.style.display = 'none';

    if (fileInput.files.length === 0) {
        responseMessage.textContent = 'Please select a file to upload.';
        responseMessage.className = 'message error';
        responseMessage.style.display = 'block';
        return;
    }

    showLoading(); // Show loading indicator

    const formData = new FormData();
    formData.append('file', fileInput.files[0]);

    try {
        const response = await fetch('/api/v1/convert', {
            method: 'POST',
            body: formData
        });

        const data = await response.json();

        if (response.ok) {
            responseMessage.textContent = data.message + (data.converted_filename ? `: ${data.converted_filename}` : '');
            responseMessage.className = 'message success';
            downloadLink.href = data.download_url;
            downloadLink.textContent = `Download ${data.converted_filename}`;
            downloadLink.style.display = 'block';
        } else {
            responseMessage.textContent = 'Error: ' + (data.error || 'Unknown error');
            responseMessage.className = 'message error';
        }
    } catch (error) {
        responseMessage.textContent = 'Network error: ' + error.message;
        responseMessage.className = 'message error';
    } finally {
        hideLoading(); // Always hide loading indicator
        responseMessage.style.display = 'block';
    }
});