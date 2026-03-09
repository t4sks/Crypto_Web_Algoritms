document.addEventListener('DOMContentLoaded', () => {
    const selectorBtn = document.getElementById('current-cipher');
    const cipherMenu = document.getElementById('cipher-menu');
    const settingsArea = document.getElementById('dynamic-setting');
    const encryptBtn = document.getElementById('encrypt-btn');

    const configs = {
        scytale: `
            <div class="setting-item">
                <span class="hint">Ключ:</span>
                <input type="number" id="cipher-key" class="dynamic-input" value="4" min="1">
            </div>

            <div class="setting-item">
                <span class="hint">Операция:</span>
                <div class="dropdown" data-dropdown>
                    <button
                        type="button"
                        class="dropdown-trigger"
                        data-target="cipher-op"
                        data-value="encrypt"
                    >
                        Зашифровать
                    </button>
                    <ul class="dropdown-menu">
                        <li class="dropdown-item" data-value="encrypt">Зашифровать</li>
                        <li class="dropdown-item" data-value="decrypt">Расшифровать</li>
                    </ul>
                </div>
                <input type="hidden" id="cipher-op" value="encrypt">
            </div>
        `,
        polybius: `
            <div class="setting-item">
                <span class="hint">Язык:</span>
                <div class="dropdown" data-dropdown>
                    <button
                        type="button"
                        class="dropdown-trigger"
                        data-target="cipher-lang"
                        data-value="russian"
                    >
                        Русский
                    </button>
                    <ul class="dropdown-menu">
                        <li class="dropdown-item" data-value="russian">Русский</li>
                        <li class="dropdown-item" data-value="english">Английский</li>
                    </ul>
                </div>
                <input type="hidden" id="cipher-lang" value="russian">
            </div>

            <div class="setting-item">
                <span class="hint">Операция:</span>
                <div class="dropdown" data-dropdown>
                    <button
                        type="button"
                        class="dropdown-trigger"
                        data-target="cipher-op"
                        data-value="encrypt"
                    >
                        Зашифровать
                    </button>
                    <ul class="dropdown-menu">
                        <li class="dropdown-item" data-value="encrypt">Зашифровать</li>
                        <li class="dropdown-item" data-value="decrypt">Расшифровать</li>
                    </ul>
                </div>
                <input type="hidden" id="cipher-op" value="encrypt">
            </div>
        `,
        caesar: `
            <div class="setting-item">
                <span class="hint">Ключ:</span>
                <input type="number" id="cipher-key" class="dynamic-input" value="2" min="1">
            </div>

            <div class="setting-item">
                <span class="hint">Операция:</span>
                <div class="dropdown" data-dropdown>
                    <button
                        type="button"
                        class="dropdown-trigger"
                        data-target="cipher-op"
                        data-value="encrypt"
                    >
                        Зашифровать
                    </button>
                    <ul class="dropdown-menu">
                        <li class="dropdown-item" data-value="encrypt">Зашифровать</li>
                        <li class="dropdown-item" data-value="decrypt">Расшифровать</li>
                    </ul>
                </div>
                <input type="hidden" id="cipher-op" value="encrypt">
            </div>
        `
    };

    const apiNames = {
        scytale: 'Scytale',
        polybius: 'Polibius',
        caesar: 'Caesar'
    };

    function closeAllMenus(exceptMenu = null) {
        document.querySelectorAll('.dropdown-menu.show').forEach(menu => {
            if (menu !== exceptMenu) {
                menu.classList.remove('show');
            }
        });
    }

    function updateActionButton() {
        const op = document.getElementById('cipher-op')?.value || 'encrypt';
        encryptBtn.textContent = op === 'decrypt' ? 'Расшифровать' : 'Зашифровать';
    }

    function initInnerDropdowns() {
        const dropdowns = settingsArea.querySelectorAll('[data-dropdown]');

        dropdowns.forEach(dropdown => {
            const trigger = dropdown.querySelector('.dropdown-trigger');
            const menu = dropdown.querySelector('.dropdown-menu');
            const items = dropdown.querySelectorAll('.dropdown-item');

            if (!trigger || !menu) {
                return;
            }

            trigger.addEventListener('click', (e) => {
                e.stopPropagation();
                closeAllMenus(menu);
                menu.classList.toggle('show');
            });

            items.forEach(item => {
                item.addEventListener('click', () => {
                    const value = item.dataset.value || '';
                    const text = item.textContent.trim();
                    const targetId = trigger.dataset.target;

                    trigger.textContent = text;
                    trigger.dataset.value = value;

                    if (targetId) {
                        const hiddenInput = document.getElementById(targetId);
                        if (hiddenInput) {
                            hiddenInput.value = value;
                        }
                    }

                    menu.classList.remove('show');

                    if (targetId === 'cipher-op') {
                        updateActionButton();
                    }
                });
            });
        });

        updateActionButton();
    }

    function setCipher(cipherKey, buttonText) {
        selectorBtn.textContent = buttonText;
        selectorBtn.dataset.apiName = apiNames[cipherKey] || '';
        settingsArea.innerHTML = configs[cipherKey] || '';

        initInnerDropdowns();
        updateActionButton();
    }

    selectorBtn.addEventListener('click', (e) => {
        e.stopPropagation();
        closeAllMenus(cipherMenu);
        cipherMenu.classList.toggle('show');
    });

    cipherMenu.querySelectorAll('.dropdown-item').forEach(item => {
        item.addEventListener('click', () => {
            const value = item.dataset.value;
            const text = item.textContent.trim();

            setCipher(value, text);
            cipherMenu.classList.remove('show');
        });
    });

    window.addEventListener('click', () => {
        closeAllMenus();
    });

    setCipher('scytale', 'Scytale');

    encryptBtn.addEventListener('click', handleExecution);
});

async function handleExecution() {
    const outputField = document.getElementById('output-text');
    const responseArea = document.getElementById('response-area');
    const responseMessage = document.getElementById('response-message');
    const requestIdField = document.getElementById('request-id');
    const encryptBtn = document.getElementById('encrypt-btn');
    const userInput = document.getElementById('user-input').value.trim();
    const currentCipher = document.getElementById('current-cipher');

    if (!userInput) {
        responseArea.style.display = 'block';
        responseMessage.textContent = 'Ошибка:';
        outputField.style.color = '#ff4444';
        outputField.textContent = 'Введите текст.';
        if (requestIdField) {
            requestIdField.textContent = '';
        }
        return;
    }

    responseArea.style.display = 'block';
    responseMessage.textContent = 'Результат:';
    outputField.style.color = '';
    outputField.textContent = 'Обработка...';

    if (requestIdField) {
        requestIdField.textContent = '';
    }

    encryptBtn.disabled = true;

    const requestData = {
        algoritm: currentCipher.dataset.apiName || '',
        data: userInput,
        language: document.getElementById('cipher-lang')?.value || '',
        operation: document.getElementById('cipher-op')?.value || 'encrypt',
        key: parseInt(document.getElementById('cipher-key')?.value, 10) || 0
    };

    try {
        const response = await fetch('/api', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(requestData)
        });

        const result = await response.json();

        if (!response.ok) {
            throw new Error(result.error || 'Ошибка сервера');
        }

        outputField.style.color = '';
        outputField.textContent = result.result || '';

        if (requestIdField) {
            requestIdField.textContent = result.request_id
                ? `Request-ID: ${result.request_id}`
                : '';
        }
    } catch (err) {
        responseMessage.textContent = 'Ошибка:';
        outputField.style.color = '#ff4444';
        outputField.textContent = err.message;

        if (requestIdField) {
            requestIdField.textContent = '';
        }
    } finally {
        encryptBtn.disabled = false;
    }
}