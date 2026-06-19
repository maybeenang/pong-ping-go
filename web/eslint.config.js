import js from '@eslint/js';
import reactHooks from 'eslint-plugin-react-hooks';
import reactRefresh from 'eslint-plugin-react-refresh';
import unusedImports from 'eslint-plugin-unused-imports';
import { defineConfig, globalIgnores } from 'eslint/config';
import globals from 'globals';
import tseslint from 'typescript-eslint';

/** @type {import('eslint').Linter.Config[]} */
export default defineConfig([
    globalIgnores(['dist']),
    {
        plugins: {
            'unused-imports': unusedImports,
        },
        files: ['**/*.{ts,tsx}'],
        rules: {
            'no-unused-vars': 'off',
            'no-console': 'warn',
            'unused-imports/no-unused-imports': 'error',
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
]);
