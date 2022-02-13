export interface Response {
    ok: boolean
    errCode: number
    errMsg: string
    data: object
    timestamp: Date
}

export interface ArticleRaw {
    title: string
    body: string
}

export interface IdValue {
    id: string,
    value: string
}