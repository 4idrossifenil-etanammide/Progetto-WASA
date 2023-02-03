import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import Photo from './components/Photo.vue'
import CommentModal from './components/CommentModal.vue'
import WASAPhotoBanner from './components/WASAPhotoBanner.vue'
import UploadPhotoModal from './components/UploadPhotoModal.vue'
import ChangeNameModal from './components/ChangeNameModal.vue'

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("Photo", Photo);
app.component("CommentModal", CommentModal);
app.component("WASAPhotoBanner", WASAPhotoBanner);
app.component("UploadPhotoModal", UploadPhotoModal);
app.component("ChangeNameModal", ChangeNameModal)
app.use(router)
app.mount('#app')
