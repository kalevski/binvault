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
    themeColor: '#9c27b0',
    name: 'BINVAULT.io',
    nameLink: '/',
    logo: '/logo.png',
    auto2top: true,
    routerMode: 'hash',
})