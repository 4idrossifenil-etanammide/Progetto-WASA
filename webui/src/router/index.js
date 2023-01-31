import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import HomePageView from '../views/HomePageView.vue'
import ProfileView from '../views/ProfileView.vue'
import PersonalPageView from '../views/PersonalPageView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView},
		{path: '/homepage', component: HomePageView},
		{path: '/homepage/profile', component: ProfileView},
		{path: '/homepage/profile/personal', component: PersonalPageView}
	]
})

export default router
