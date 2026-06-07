import js from '@eslint/js'
import globals from 'globals'
import reactHooks from 'eslint-plugin-react-hooks'
import reactRefresh from 'eslint-plugin-react-refresh'
import tseslint from 'typescript-eslint'
import { defineConfig, globalIgnores } from 'eslint/config'
import prettier from 'eslint-config-prettier/flat';

/** @type {import('eslint').Linter.Config[]} */
export default defineConfig([
    globalIgnores(['dist']),
    {
        plugins: [
            "unused-imports"
        ],
        files: ['**/*.{ts,tsx}'],
        rules: {
            "no-unused-vars": "off",
            "unused-imports/no-unused-imports": "error"
        },
        extends: [
            js.configs.recommended,
            tseslint.configs.recommended,
            reactHooks.configs.flat.recommended,
            reactRefresh.configs.vite,
        ],
        languageOptions: {
            globals: globals.browser,
        },
    },
    prettier,
])
