window.buildRequestData = function () {
    const currentCipher = document.getElementById('current-cipher');
    const cipherKey = currentCipher.dataset.cipherKey || '';

    const base = {
        algorithm: currentCipher.dataset.apiName || '',
        data: document.getElementById('user-input').value.trim(),
        language: document.getElementById('cipher-lang')?.value || '',
        operation: document.getElementById('cipher-op')?.value || 'encrypt',
    };

    if (cipherKey === 'cardano') {
        base.code = (document.getElementById('cardano-code')?.value || '').trim();
    }

    if (cipherKey === 'gronsfeld') {
        base.keyString = (document.getElementById('cipher-key')?.value || '').trim();
    } else {
        base.key = parseInt(document.getElementById('cipher-key')?.value, 10) || 0;
    }

    return base;
};

window.handleExecution = async function () {
    const outputField = document.getElementById('output-text');
    const responseArea = document.getElementById('response-area');
    const responseMessage = document.getElementById('response-message');
    const requestIdField = document.getElementById('request-id');
    const statusField = document.getElementById('response-status');
    const codeWrap = document.getElementById('cardano-code-wrap');
    const codeOutput = document.getElementById('cardano-code-output');
    const actionBtn = document.getElementById('encrypt-btn');
    const currentCipher = document.getElementById('current-cipher');
    const cipherKey = currentCipher.dataset.cipherKey || '';

    const requestData = window.buildRequestData();

    responseArea.classList.remove('is-hidden');
    responseMessage.textContent = 'Результат:';
    outputField.classList.remove('text-error');
    outputField.textContent = 'Обработка...';
    requestIdField.textContent = '';
    statusField.textContent = '';
    codeWrap.classList.add('is-hidden');
    codeOutput.textContent = '';

    actionBtn.disabled = true;

    try {
        const response = await fetch('/api', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(requestData)
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
            throw new Error(result.error || 'Ошибка сервера');
        }

        outputField.classList.remove('text-error');
        outputField.textContent = result.result || '';

        if (
            cipherKey === 'cardano' &&
            requestData.operation === 'encrypt' &&
            result.cardano_code
        ) {
            codeWrap.classList.remove('is-hidden');
            codeOutput.textContent = result.cardano_code;
        }
    } catch (err) {
        responseMessage.textContent = 'Ошибка:';
        outputField.classList.add('text-error');
        outputField.textContent = err.message || 'Неизвестная ошибка';
        codeWrap.classList.add('is-hidden');
        codeOutput.textContent = '';
    } finally {
        actionBtn.disabled = false;
    }
};