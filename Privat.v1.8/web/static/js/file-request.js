document.addEventListener('DOMContentLoaded', () => {
    const MAX_FILE_SIZE = 64 * 1024;
    const ALLOWED_EXTENSIONS = ['txt'];

    const fileInput = document.getElementById('source-file');
    const fileDropzone = document.getElementById('file-dropzone');

    const fileNameField = document.getElementById('file-name');
    const fileSizeField = document.getElementById('file-size');
    const fileStatusField = document.getElementById('file-status');
    const previewField = document.getElementById('file-preview');
    const downloadBtn = document.getElementById('download-result-btn');

    if (!fileInput || !fileDropzone || !downloadBtn) {
        return;
    }

    let currentFile = null;
    let currentPreviewText = '';

    function safeText(value) {
        return String(value || '');
    }

    function getExtension(filename) {
        const parts = String(filename || '').toLowerCase().split('.');
        if (parts.length < 2) {
            return '';
        }
        return parts.pop();
    }

    function hasDoubleExtension(filename) {
        const parts = String(filename || '').toLowerCase().split('.');
        if (parts.length < 3) {
            return false;
        }

        const suspicious = ['php', 'js', 'exe', 'sh', 'bat', 'cmd', 'html'];
        return parts.slice(1, -1).some(part => suspicious.includes(part));
    }

    function formatBytes(bytes) {
        if (bytes < 1024) {
            return `${bytes} Б`;
        }
        if (bytes < 1024 * 1024) {
            return `${(bytes / 1024).toFixed(1)} КБ`;
        }
        return `${(bytes / (1024 * 1024)).toFixed(1)} МБ`;
    }

    function validateFile(file) {
        const ext = getExtension(file.name);

        if (!ext || !ALLOWED_EXTENSIONS.includes(ext)) {
            return 'Неподдерживаемый тип файла';
        }

        if (hasDoubleExtension(file.name)) {
            return 'Подозрительное двойное расширение';
        }

        if (file.size <= 0) {
            return 'Пустой файл';
        }

        if (file.size > MAX_FILE_SIZE) {
            return 'Файл слишком большой';
        }

        return '';
    }

    function clearDownloadLink() {
        downloadBtn.href = '#';
        downloadBtn.removeAttribute('download');
        downloadBtn.classList.add('is-disabled');
        downloadBtn.setAttribute('aria-disabled', 'true');
    }

    function setupDownloadLink(downloadUrl) {
        if (!downloadUrl) {
            clearDownloadLink();
            return;
        }

        downloadBtn.href = downloadUrl;
        downloadBtn.removeAttribute('download');
        downloadBtn.classList.remove('is-disabled');
        downloadBtn.setAttribute('aria-disabled', 'false');
    }

    function resetFileState() {
        currentFile = null;
        currentPreviewText = '';
        fileInput.value = '';

        fileNameField.textContent = '—';
        fileSizeField.textContent = '—';
        fileStatusField.textContent = 'Файл не выбран';
        previewField.textContent = 'Здесь появится содержимое файла.';

        clearDownloadLink();
    }

    function showError(message) {
        const outputField = document.getElementById('output-text');
        const responseArea = document.getElementById('response-area');
        const responseMessage = document.getElementById('response-message');
        const requestIdField = document.getElementById('request-id');
        const statusField = document.getElementById('response-status');
        const codeWrap = document.getElementById('cardano-code-wrap');
        const codeOutput = document.getElementById('cardano-code-output');

        responseArea.classList.remove('is-hidden');
        responseMessage.textContent = 'Ошибка:';
        outputField.classList.add('text-error');
        outputField.textContent = message;
        requestIdField.textContent = '';
        statusField.textContent = '';
        codeWrap.classList.add('is-hidden');
        codeOutput.textContent = '';

        clearDownloadLink();
    }

    function handleFileSelection(file) {
        clearDownloadLink();

        const validationError = validateFile(file);
        if (validationError) {
            currentFile = null;
            currentPreviewText = '';

            fileNameField.textContent = safeText(file.name || '—');
            fileSizeField.textContent = formatBytes(file.size || 0);
            fileStatusField.textContent = validationError;
            previewField.textContent = 'Файл отклонён.';
            showError(validationError);
            return;
        }

        currentFile = file;
        fileNameField.textContent = safeText(file.name);
        fileSizeField.textContent = formatBytes(file.size);
        fileStatusField.textContent = 'Файл загружен';

        const reader = new FileReader();
        reader.onload = () => {
            currentPreviewText = typeof reader.result === 'string' ? reader.result : '';
            previewField.textContent = currentPreviewText || 'Файл пустой.';
        };
        reader.onerror = () => {
            currentPreviewText = '';
            previewField.textContent = 'Не удалось прочитать файл.';
            fileStatusField.textContent = 'Ошибка чтения';
            clearDownloadLink();
        };

        reader.readAsText(file);
    }

    function buildFileFormData() {
        const currentCipher = document.getElementById('current-cipher');
        const cipherKey = currentCipher?.dataset?.cipherKey || '';
        const formData = new FormData();

        formData.append('file', currentFile);
        formData.append('algorithm', currentCipher?.dataset?.apiName || '');
        formData.append('language', document.getElementById('cipher-lang')?.value || '');
        formData.append('operation', document.getElementById('cipher-op')?.value || 'encrypt');

        if (cipherKey === 'cardano') {
            formData.append('code', (document.getElementById('cardano-code')?.value || '').trim());
        }

        if (cipherKey === 'gronsfeld' || cipherKey === 'vigenere') {
            formData.append('keyString', (document.getElementById('cipher-key')?.value || '').trim());
        } else {
            formData.append('key', String(parseInt(document.getElementById('cipher-key')?.value, 10) || 0));
        }

        return formData;
    }

    window.handleFileExecution = async function () {
        const outputField = document.getElementById('output-text');
        const responseArea = document.getElementById('response-area');
        const responseMessage = document.getElementById('response-message');
        const requestIdField = document.getElementById('request-id');
        const statusField = document.getElementById('response-status');
        const codeWrap = document.getElementById('cardano-code-wrap');
        const codeOutput = document.getElementById('cardano-code-output');
        const actionBtn = document.getElementById('encrypt-btn');
        const currentCipher = document.getElementById('current-cipher');
        const cipherKey = currentCipher?.dataset?.cipherKey || '';

        if (!currentFile) {
            showError('Сначала выбери файл');
            return;
        }

        responseArea.classList.remove('is-hidden');
        responseMessage.textContent = 'Результат:';
        outputField.classList.remove('text-error');
        outputField.textContent = 'Обработка...';
        requestIdField.textContent = '';
        statusField.textContent = '';
        codeWrap.classList.add('is-hidden');
        codeOutput.textContent = '';
        actionBtn.disabled = true;
        fileStatusField.textContent = 'Идёт обработка';
        clearDownloadLink();

        const formData = buildFileFormData();

        try {
            const response = await fetch('/api/file/process', {
                method: 'POST',
                body: formData
            });

            let result = {};
            try {
                result = await response.json();
            } catch {
                result = {};
            }

            statusField.textContent = `HTTP: ${response.status}`;

            const requestId = result.request_id || response.headers.get('X-Request-ID') || '';
            requestIdField.textContent = requestId ? `Request-ID: ${requestId}` : '';

            if (!response.ok) {
                throw new Error(result.error || 'Ошибка обработки файла');
            }

            outputField.classList.remove('text-error');
            outputField.textContent = result.result || '';
            fileStatusField.textContent = 'Обработано';

            if (
                cipherKey === 'cardano' &&
                formData.get('operation') === 'encrypt' &&
                result.cardano_code
            ) {
                codeWrap.classList.remove('is-hidden');
                codeOutput.textContent = result.cardano_code;
            }

            setupDownloadLink(result.download_url || '');
        } catch (err) {
            responseMessage.textContent = 'Ошибка:';
            outputField.classList.add('text-error');
            outputField.textContent = err.message || 'Неизвестная ошибка';
            codeWrap.classList.add('is-hidden');
            codeOutput.textContent = '';
            fileStatusField.textContent = 'Ошибка обработки';
            clearDownloadLink();
        } finally {
            actionBtn.disabled = false;
        }
    };

    fileInput.addEventListener('change', () => {
        const file = fileInput.files?.[0];
        if (!file) {
            resetFileState();
            return;
        }

        handleFileSelection(file);
    });

    fileDropzone.addEventListener('dragover', (e) => {
        e.preventDefault();
        fileDropzone.classList.add('dropzone-active');
    });

    fileDropzone.addEventListener('dragleave', () => {
        fileDropzone.classList.remove('dropzone-active');
    });

    fileDropzone.addEventListener('drop', (e) => {
        e.preventDefault();
        fileDropzone.classList.remove('dropzone-active');

        const file = e.dataTransfer?.files?.[0];
        if (!file) {
            return;
        }

        fileInput.files = e.dataTransfer.files;
        handleFileSelection(file);
    });

    downloadBtn.addEventListener('click', (e) => {
        if (downloadBtn.getAttribute('aria-disabled') === 'true') {
            e.preventDefault();
        }
    });

    resetFileState();
});