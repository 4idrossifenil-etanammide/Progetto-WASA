<script>

export default {
    data: function() {
        return{
            photoURL: __API_URL__ + "/profiles/" + localStorage.getItem("name") + "/photos/",
            photos: [],
            userToSearch: ''
        }
    },
    async mounted() {
        await this.caricaFoto();
    },
    methods: {
        async caricaFoto() {
            try{
                const response = await this.$axios.get("/profiles/" + localStorage.getItem("name") + "/photos")
                this.photos = response.data["photos"]
            } catch (e) {}
        },
        search() {
            this.userToSearch = this.$refs.userToSearch.value;
            localStorage.setItem("profile", this.userToSearch)
            this.$router.push("/homepage/profile");
        }
    }
}
</script>

<template>  
    <div class="page">
        <div class="header">
            <h1 class="title">WasaPHOTO</h1>
            <form class="search-form" @submit.prevent="search">
				<input class="input-form" v-model="userToSearch" ref="userToSearch" type="text" placeholder="Search username..." required>
			</form>
        </div>
        <div class="photo-container">
            <div class="photo" v-if="this.photos != null" v-for="photo in this.photos" :key="photo.photoID">
                <Photo :id="photo.photoID" :name="photo.name" :likes="photo.likeNumber" :comments="photo.commentNumber" :date="photo.date" :textComments="photo.comments"/>
            </div>
        </div>
    </div>
</template>

<style>

.input-form {
    border: 1px solid #ccc;
	border-radius: 4px; /* Aggiunge spigoli arrotondati */
    height: 30px;
}

.search-form {
    position: relative;
    left: 650px;
    top: 10px;
}

.page {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.header {
    background-color: rgb(215, 255, 255);
    height: 50px;
    width: 100%;
    display: flex;
    justify-content: center;
    margin-bottom: 10px;
}

.title{
    position: relative;
    top: 7px;
    left: 90px;
}

.photo-container{
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    overflow: auto;
}

.photo{
    margin-bottom: 15px;
}
</style>