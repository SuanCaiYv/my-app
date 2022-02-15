import {ListResult} from "../common/interface";

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

export const toListResult = function (data: object): ListResult {
    // @ts-ignore
    return new ListResultClass(data.count, data.end_page, data.list, data.next_page_num, data.page_num, data.page_size, data.total)
}