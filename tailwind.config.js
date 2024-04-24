/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./internal/view/**/*.{templ,go}"],
	theme: {
		extend: {},
	},
	plugins: [require("daisyui")],
	daisyui: {
		themes: [
			{
				dracula: {
					...require("daisyui/src/theming/themes")["dracula"],
					primary: "#BD93F9",
					secondary: "#FF79C6",
					error: "#B80C09"
				},
			},
		],
	},
};
