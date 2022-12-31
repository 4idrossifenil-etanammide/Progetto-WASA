import axios from "axios";
import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import HomePageView from '../views/HomePageView.vue'

axios.interceptors.request.use(function (config){
		config.headers.Authorization = localStorage.getItem("token");
		return config;
	}, function (error) {
		return Promise.reject(error);
	}
);

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView},
		{path: '/homepage', component: HomePageView},
	]
})

export default router
