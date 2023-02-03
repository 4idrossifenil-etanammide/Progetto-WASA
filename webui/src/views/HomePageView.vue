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
        <WASAPhotoBanner></WASAPhotoBanner>
        <div class="no-photo-label" v-show="this.photos == null">
            <h1 class="no-photo-text"> No new Photos </h1>
        </div>
        <div class="photo-container">
            <div class="photo" v-show="this.photos != null" v-for="photo in this.photos" :key="photo.photoID">
                <Photo :id="photo.photoID" :name="photo.name" :likes="photo.likeNumber" :comments="photo.commentNumber" :date="photo.date" :textComments="photo.comments"/>
            </div>
        </div>
    </div>
</template>

<style>

.no-photo-label {
    font-weight: bold;
    font-size: 20px;
    color: red;
}

.page {
    display: flex;
    flex-direction: column;
    align-items: center;
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