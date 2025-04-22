import Docsify from './Docsify'

const docsify = new Docsify()
docsify.initialize({
    homepage: 'home.md',
    coverpage: false,
    basePath: '/docs',
    loadSidebar: true,
    subMaxLevel: 2,
    loadNavbar: true,
    ga: 'G-8Q40653GG4',
    themeColor: '#9c27b0',
    name: 'BINVAULT.io',
    nameLink: 'https://binvault.io/',
    logo: '/logo.png',
})