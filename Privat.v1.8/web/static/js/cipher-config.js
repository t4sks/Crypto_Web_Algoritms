window.configs = {
    scytale: `
        <div class="setting-item">
            <span class="hint">Ключ:</span>
            <input type="number" id="cipher-key" class="dynamic-input" value="4" min="1">
        </div>

        <div class="setting-item">
            <span class="hint">Операция:</span>
            <div class="dropdown" data-dropdown>
                <button type="button" class="dropdown-trigger" data-target="cipher-op" data-value="encrypt">
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
                <button type="button" class="dropdown-trigger" data-target="cipher-lang" data-value="russian">
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
                <button type="button" class="dropdown-trigger" data-target="cipher-op" data-value="encrypt">
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
                <button type="button" class="dropdown-trigger" data-target="cipher-op" data-value="encrypt">
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
    cardano: `
        <div class="setting-item">
            <span class="hint">Размер k:</span>
            <input type="number" id="cipher-key" class="dynamic-input" value="2" min="1">
            <div class="field-hint">Большая таблица будет размером 2k × 2k</div>
        </div>

        <div class="setting-item">
            <span class="hint">Операция:</span>
            <div class="dropdown" data-dropdown>
                <button type="button" class="dropdown-trigger" data-target="cipher-op" data-value="encrypt">
                    Зашифровать
                </button>
                <ul class="dropdown-menu">
                    <li class="dropdown-item" data-value="encrypt">Зашифровать</li>
                    <li class="dropdown-item" data-value="decrypt">Расшифровать</li>
                </ul>
            </div>
            <input type="hidden" id="cipher-op" value="encrypt">
        </div>

        <div class="setting-item is-hidden" id="cardano-code-group">
            <span class="hint">Код решётки:</span>
            <input type="text" id="cardano-code" class="dynamic-input" placeholder="Например: 2413">
            <div class="field-hint">Для расшифрования вставь код, который был выдан при шифровании</div>
        </div>
    `,
    gronsfeld: `
        <div class="setting-item">
            <span class="hint">Ключ:</span>
            <input type="text" id="cipher-key" class="dynamic-input" value="2015" inputmode="numeric" placeholder="Пример: 2015">
        </div>

        <div class="setting-item">
            <span class="hint">Операция:</span>
            <div class="dropdown" data-dropdown>
                <button type="button" class="dropdown-trigger" data-target="cipher-op" data-value="encrypt">
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

window.apiNames = {
    scytale: 'Scytale',
    polybius: 'Polibius',
    caesar: 'Caesar',
    cardano: 'Cardano',
    gronsfeld: 'Gronsfeld',
};