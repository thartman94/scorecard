import tseslint from "typescript-eslint";
import pluginReact from "eslint-plugin-react";
import importPlugin from "eslint-plugin-import";

export default [
    { files: ["**/*.{js,mjs,cjs,ts,jsx,tsx}"] },
    ...tseslint.configs.recommended,
    pluginReact.configs.flat.recommended,
    importPlugin.flatConfigs.recommended,
    importPlugin.flatConfigs.typescript,
    {
        ignores: ["**/node_modules", "eslint.config.js"],
    },
    {
        rules: {
            // React > 17 does not require react to be in scope for JSX
            "react/react-in-jsx-scope": "off",

            // no console.logs or warnings
            "no-console": ["warn", { allow: ["error"] }],

            // Sort imports using import plugin
            "sort-imports": "off", // disable default eslint sorting
            "import/first": "error",
            "import/newline-after-import": "error",
            "import/no-duplicates": "error",
            "import/order": [
                "error",
                {
                    alphabetize: { order: "asc" },
                    "newlines-between": "always",
                },
            ],
        },
    },
    {
        settings: {
            react: {
                version: "detect",
            },
        },
    },
];
