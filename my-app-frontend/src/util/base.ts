import {ArticleRaw, Catalog, IdName, ListResult} from "../common/interface";

export const parseBoolean = function (str: string): boolean {
    if (str === "false") {
        return false
    } else if (str === "true") {
        return true
    } else if (str === "") {
        return false
    } else {
        throw new Error("fuck input")
    }
}

class ListResultClass implements ListResult {
    count: number;
    endPage: boolean;
    list: Array<object>;
    nextPageNum: number;
    pageNum: number;
    pageSize: number;
    total: number;

    constructor(count: number, endPage: boolean, list: Array<object>, nextPageNum: number, pageNum: number, pageSize: number, total: number) {
        this.count = count;
        this.endPage = endPage;
        this.list = list;
        this.nextPageNum = nextPageNum;
        this.pageNum = pageNum;
        this.pageSize = pageSize;
        this.total = total;
    }
}

class ArticleRawClass implements ArticleRaw {
    articleId: string;
    articleName: string;
    author: string;
    catalog: Catalog;
    content: string;
    coverImg: string;
    kind: IdName;
    releaseTime: Date;
    summary: string;
    tagList: Array<IdName>;
    visibility: number;

    constructor(articleId: string, articleName: string, author: string, catalog: Catalog, content: string, coverImg: string, kind: IdName, releaseTime: Date, summary: string, tagList: Array<IdName>, visibility: number) {
        this.articleId = articleId;
        this.articleName = articleName;
        this.author = author;
        this.catalog = catalog;
        this.content = content;
        this.coverImg = coverImg;
        this.kind = kind;
        this.releaseTime = releaseTime;
        this.summary = summary;
        this.tagList = tagList;
        this.visibility = visibility;
    }
}

class IdNameClass implements IdName {
    id: string;
    name: string;

    constructor(id: string, name: string) {
        this.id = id;
        this.name = name;
    }
}

class CatalogClass implements Catalog {
    name: string;
    children: Array<Catalog>;

    constructor(name: string, children: Array<Catalog>) {
        this.children = children;
        this.name = name;
    }
}

export const toListResult = function (data: object): ListResult {
    // @ts-ignore
    return new ListResultClass(data.count, data.end_page, data.list, data.next_page_num, data.page_num, data.page_size, data.total)
}

export const toArticleRaw = function (data: string): ArticleRaw {
    return toArticleRawWithObject(JSON.parse(data))
}

export const toArticleRawWithObject = function (obj: any): ArticleRaw {
    let tagList: Array<IdName> = []
    for (let i of obj.tag_list) {
        tagList.push(new IdNameClass(i.tag_id, i.tag_name))
    }
    return new ArticleRawClass(obj.article_id, obj.article_name, obj.author, catalogFunc(obj.catalog), obj.content, obj.cover_img, new IdNameClass(obj.kind.kind_id, obj.kind.kind_name), obj.release_time, obj.summary, tagList, obj.visibility)
}

const catalogFunc = function (obj: any): Catalog {
    if (obj.children.length === 0) {
        return new CatalogClass(obj.name, [])
    } else {
        let result = new CatalogClass(obj.name, [])
        for (let i of obj.children) {
            let t = catalogFunc(i)
            result.children.push(t)
        }
        return result
    }
}