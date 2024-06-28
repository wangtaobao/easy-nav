import request from "@/api/request"

export function getIndexList(params?: object) {
    return request({
        url: "/api",
        method: "get",
        params
    })
}

export function addUrl(data: object) {
    return request({
        url: "/api/add",
        method: "post",
        data
    })
}

export function delUrl(id) {
    return request({
        url: "/api/edit/" + id,
        method: "delete",
    })
}

export function updateUrl(data: object) {
    return request({
        url: "/api/edit",
        method: "put",
        data
    })
}