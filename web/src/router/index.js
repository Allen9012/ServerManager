import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [{
        path: '/',
        redirect: '/login'
    },
    {
        path: '/init',
        name: 'Init',
        component: () =>
            import ('@/view/init/index.vue')
    },
    {
        path: '/login',
        name: 'Login',
        component: () =>
            import ('@/view/login/index.vue')
    },
    {
        path: '/:catchAll(.*)',
        meta: {
            closeTab: true,
        },
        component: () =>
            import ('@/view/error/index.vue')
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    strict: false,
    scrollBehavior: () => ({ left: 0, top: 0 }),
    routes,
})

export default router