// import Router from 'vue-router'
import { createRouter, createWebHashHistory } from "vue-router"

// import NProgress from 'nprogress'
// import 'nprogress/nprogress.css'

import NProgress from "@/config/nprogress"

// const router = new Router({
//   routes: [
//     // your routes here
//   ]
// })

const routers = [
  {
    path: '/',
    meta: {
      title: "导航"
    },
    component: () => import("@/views/Index.vue")
  },
  {
    path: '/ws',
    meta: {
      title: "ws"
    },
    component: () => import("@/views/Ws.vue")
  },
  // {
  //   path: '/edit',
  //   meta: {
  //     title: "edit"
  //   },
  //   component: () => import("@/views/EditUrl.vue")
  // },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes: routers
})

router.beforeEach((to, from, next) => {
  NProgress.start()
  next()
})

router.afterEach(() => {
  NProgress.done()
})

router.onError(error => {
  NProgress.done()
  console.warn(error.message)
})

export default router