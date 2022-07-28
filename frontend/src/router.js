
import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router);


import PetManager from "./components/PetManager"

import ItemManager from "./components/ItemManager"
import OrderItemManager from "./components/OrderItemManager"
import CustomerManager from "./components/CustomerManager"

export default new Router({
    // mode: 'history',
    base: process.env.BASE_URL,
    routes: [
            {
                path: '/pets',
                name: 'PetManager',
                component: PetManager
            },

            {
                path: '/items',
                name: 'ItemManager',
                component: ItemManager
            },
            {
                path: '/orderItems',
                name: 'OrderItemManager',
                component: OrderItemManager
            },
            {
                path: '/customers',
                name: 'CustomerManager',
                component: CustomerManager
            },



    ]
})
