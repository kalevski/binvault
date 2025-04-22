export type DocsifyConfig = {
    repo: string
    maxLevel: number
    coverpage: boolean
    renderer: {
        [key: string]: (code: string, lang: string) => void
    },
    alias: {
        [key: string]: string
    },
    auto2top: boolean
    autoHeader: boolean
    basePath: string
    cornerExternalLinkTarget: string,
    el: string
    homepage: string
    logo: string
    onlyCover: boolean
    loadSidebar: boolean
    subMaxLevel: number
    loadNavbar: boolean
    ga: string
    themeColor: string
    name: string
    nameLink: string
}

class Docsify {

    initialize(config: Partial<DocsifyConfig>) {
        window['$docsify'] = config
    }

}

export default Docsify