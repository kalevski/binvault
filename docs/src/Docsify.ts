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
    routerMode: 'hash' | 'history',
}

export interface DocsifyPlugin {
    start(): void
    dispose(): void
}

class Docsify {

    plugins: DocsifyPlugin[] = []

    initialize(config: Partial<DocsifyConfig>) {
        window['$docsify'] = config
        for (const plugin of this.plugins) {
            plugin.start()
        }
    }

    addPlugin(plugin: DocsifyPlugin) {
        this.plugins.push(plugin)
    }

}

export default Docsify