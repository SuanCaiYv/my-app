export interface Response {
    ok: boolean
    errCode: number
    errMsg: string
    data: object
    timestamp: Date
}

export interface ArticleLiteRaw {
    articleId: string
    articleName: string
    summary: string
    visibility: number
}

export interface Catalog {
    name: string
    children: Array<Catalog>
}

export interface ArticleRaw {
    articleId: string
    articleName: string
    author: string
    summary: string
    coverImg: string
    content: string
    catalog: Catalog,
    kind: IdName
    tagList: Array<IdName>
    releaseTime: Date
    visibility: number
}

export interface IdName {
    id: string,
    name: string
}

export interface ListResult {
    total: number,
    count: number,
    pageNum: number,
    pageSize: number,
    nextPageNum: number,
    endPage: boolean,
    list: Array<any>,
}