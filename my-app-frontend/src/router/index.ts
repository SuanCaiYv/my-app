import {createRouter, createWebHistory, Router, RouteRecordRaw} from "vue-router";

const routes: Array<RouteRecordRaw> = [
    {
        path: "/home",
        alias: "/",
        name: "home",
        meta: {
            title: "首页"
        },
        component: () => import("../views/Home.vue")
    },
    {
        path: "/about",
        alias: "/me",
        name: "about",
        meta: {
            title: "关于我"
        },
        component: () => import("../views/About.vue")
    },
    {
        path: "/t",
        name: "test",
        component: () => import("../views/Test.vue")
    },
    {
        path: "/admin",
        name: "admin",
        component: () => import("../views/Admin.vue")
    },
    {
        path: "/sign",
        name: "sign",
        component: () => import("../views/Sign.vue")
    },
    // 文章管理
    {
        path: "/article_list",
        name: "article_list",
        component: () => import("../views/ArticleList.vue")
    },
    // 标签页和分类页管理
    {
        path: "/state_manager",
        name: "state_manager",
        component: () => import("../views/StateManager.vue")
    },
    // 完整文章
    {
        path: "/display",
        name: "display",
        component: () => import("../views/Display.vue")
    },
    // 清空缓存
    {
        path: "/truncate",
        name: "truncate",
        component: () => import("../views/Truncate.vue")
    },
    // 编辑器
    {
        path: "/editor",
        name: "editor",
        component: () => import("../views/Editor.vue")
    }
]

const router: Router = createRouter({
    history: createWebHistory(),
    routes: routes
})

export default router