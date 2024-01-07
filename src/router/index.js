import {createRouter, createWebHistory} from "vue-router";
import Main from "@/components/Main.vue";

const routes = [
    {
        path: "/",
        redirect: "/main",
    },
    {
        path: "/main",
        name: "Main",
        component: Main,
    }
]

const router = new createRouter({
    history: createWebHistory(),
    routes
})

export default router