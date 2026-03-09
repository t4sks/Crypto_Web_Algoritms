document.addEventListener('DOMContentLoaded', () => {
    const selectorBtn = document.getElementById('current-cipher');
    const menu = document.getElementById('cipher-menu');
    const settingsArea = document.getElementById('dynamic-setting');
    const menuItems = document.querySelectorAll('.cipher-menu li');
    const encryptBtn = document.getElementById('encrypt-btn');

    // Шаблоны настроек с улучшенной структурой классов для CSS
    const configs = {
        skitala: `
            <div class="setting-item">
                <span class="hint">Ключ:</span>
                <input type="number" id="cipher-key" class="dynamic-input" value="4" min="2">
            </div>
            <div class="setting-item">
                <span class="hint">Операция:</span>
                <select id="cipher-op" class="dynamic-select">
                    <option value="encrypt">Зашифровать</option>
                    <option value="decrypt">Расшифровать</option>
                </select>
            </div>`,
        polybius: `
            <div class="setting-item">
                <span class="hint">Язык:</span>
                <select id="cipher-lang" class="dynamic-select">
                    <option value="russian">Русский</option>
                    <option value="english">Англиский</option>
                </select>
            </div>
            <div class="setting-item">
                <span class="hint">Операция:</span>
                <select id="cipher-op" class="dynamic-select">
                    <option value="encrypt">Зашифровать</option>
                    <option value="decrypt">Расшифровать</option>
                </select>
            </div>`
    };

    // 1. Управление меню
    selectorBtn.addEventListener('click', (e) => {
        e.stopPropagation();
        menu.classList.toggle('show');
    });

    // 2. Смена шифра
    menuItems.forEach(item => {
        item.addEventListener('click', () => {
            const val = item.getAttribute('data-value');
            // Маппинг для соответствия вашему Go API
            const apiName = val === 'skitala' ? 'Scytale' : 'Polibius';

            selectorBtn.innerText = item.innerText;
            selectorBtn.setAttribute('data-api-name', apiName);
            settingsArea.innerHTML = configs[val] || '';
            menu.classList.remove('show');
        });
    });

    window.addEventListener('click', () => menu.classList.remove('show'));

    // Инициализация по умолчанию
    selectorBtn.setAttribute('data-api-name', 'Scytale');
    settingsArea.innerHTML = configs.skitala;

    // 3. Вызов API
    encryptBtn.addEventListener('click', handleExecution);
});

async function handleExecution() {
    const outputField = document.getElementById('output-text');
    const responseArea = document.getElementById('response-area');
    const userInput = document.getElementById('user-input').value;
    const currentCipher = document.getElementById('current-cipher');

    if (!userInput) {
        alert("Введите текст!");
        return;
    }

    responseArea.style.display = "block";
    outputField.style.color = "inherit";
    outputField.textContent = "Processing...";

    // БЕЗОПАСНЫЙ СБОР ДАННЫХ: проверяем существование элементов перед чтением [cite: 3, 6]
    const requestData = {
        algoritm: currentCipher.getAttribute('data-api-name'),
        data: userInput,
        language: document.getElementById('cipher-lang')?.value || "",
        operation: document.getElementById('cipher-op')?.value || "encrypt",
        key: parseInt(document.getElementById('cipher-key')?.value) || 0
    };

    try {
        const response = await fetch('/api', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(requestData)
        });

        const result = await response.json();

        if (!response.ok) {
            throw new Error(result.error||"Ошибка сервера");
        } else {
            outputField.style.color = "inherit";
            outputField.textContent = result.result;
        }
    } catch (err) {
        outputField.style.color = "#ff4444";
        outputField.textContent = "Error " + err.message;
    }
}