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
    }
]

const router: Router = createRouter({
    history: createWebHistory(),
    routes: routes
})

export default router