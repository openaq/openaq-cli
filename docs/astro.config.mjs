import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

// https://astro.build/config
export default defineConfig({
	integrations: [
		starlight({
			title: 'OpenAQ CLI',
			favicon: '/favicon.svg',
			  head: [
				{
				tag: 'link',
				attrs: {
					rel: 'icon',
					href:'/favicon.ico',
					sizes: '32x32',"128x128","180x180","192x192"
				},
				},
			],
			logo: {
				light: './src/assets/logo-light.svg',
				dark: './src/assets/logo-dark.svg',
			},
			customCss: [
				'./src/styles/openaq.css',
			],
			social: {
				github: 'https://github.com/openaq/openaq-cli',
			},
			sidebar: [
				{
					label: 'Start here',
					autogenerate: { directory: 'guides' },

				},
				{
					label: 'How-to guides',
					autogenerate: { directory: 'how-to' },
				},
				{
					label: 'Reference',
					autogenerate: { directory: 'reference' },
				},
			],
			components: {
				Hero: './src/components/Hero.astro',
			},
		}),
	],
});
