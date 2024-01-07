import {createRouter, createWebHistory} from "vue-router";
import Main from "@/components/Main.vue";
import Upload from "@/components/Upload.vue";
import Nofound from "@/components/nofound.vue";
import Find from "@/components/Find.vue";

const routes = [
    {
        path: "/",
        redirect: "/main",
    },
    {
        path: "/main",
        name: "Main",
        component: Main,
    },
    {
        path: "/upload",
        name: "Upload",
        component: Upload,
    },
    {
        path: "/find",
        name: "Find",
        component: Find,
    },
    {
        name: "/noFound",
        path: "/:pathMatch(.*)",
        component: Nofound,
    },
]

const router = new createRouter({
    history: createWebHistory(),
    routes
})

export default router