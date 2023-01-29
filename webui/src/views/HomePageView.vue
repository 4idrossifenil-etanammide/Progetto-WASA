<script>

export default {
    data: function() {
        return{
            photoURL: __API_URL__ + "/profiles/" + localStorage.getItem("name") + "/photos/",
            photos: []
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
    }
}
</script>

<template>  
    <div class="page">
        <div class="header">
            <h1 class="title">WasaPHOTO</h1>
        </div>
        <div class="photo-container">
            <div class="photo" v-if="this.photos.length != 0" v-for="photo in this.photos">
                <Photo :id="photo.photoID" :name="photo.name" :likes="photo.likeNumber" :comments="photo.commentNumber" :date="photo.date" :textComments="photo.comments"/>
            </div>
        </div>
    </div>
</template>

<style>

.page {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.header {
    height: 50px;
    display: flex;
    justify-content: center;
}

.title {
    text-align: center;
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