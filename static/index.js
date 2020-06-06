const Board = () => import('./components/Board.js')
const Home = () => import('./components/Home.js')
const CreateBoardForm = () => import('./components/CreateBoardForm.js')
const LayoutDefault = () => import('./layouts/LayoutDefault.js')

const routes = [
  {
    path: '/',
    component: LayoutDefault,
    children: [
      { name: 'home', path: '', component: Home },
      { name: 'create-board', path: '/board', component: CreateBoardForm },
      { name: 'board', path: '/board/:id', component: Board }
    ]
  }
]

const router = new VueRouter({
  routes
})

new Vue({
  router
}).$mount('#app')
