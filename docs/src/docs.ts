import Docsify from './Docsify'

const docsify = new Docsify()
docsify.initialize({
    homepage: 'home.md',
    coverpage: false,
    basePath: '/docs',
    loadSidebar: true,
    subMaxLevel: 1,
    loadNavbar: true,
    ga: 'G-8Q40653GG4',
    themeColor: '#e2516f',
    name: 'BINVAULT.io',
    nameLink: 'https://binvault.io/',
    logo: '/logo.png',
})