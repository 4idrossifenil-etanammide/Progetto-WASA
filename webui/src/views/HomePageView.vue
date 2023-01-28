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
                <Photo :id="photo.photoID" :name="photo.name" :likes="photo.likeNumber" :comments="photo.commentNumber" :date="photo.date"/>
            </div>
        </div>
        <div class="footer">
            <button>Bottone 1</button>
            <button>Bottone 2</button>
            <button>Aggiungi Foto</button>
        </div>
    </div>
</template>

<style>

.page {
    display: flex;
    flex-direction: column;
    align-items: center;
    height: 100vh; /* vh è un unità di misura pari al 1% dell'altezza della pagina */
}

.header {
    width: 100%;
    display: flex;
    justify-content: center;
}

.title {
    text-align: center;
}

.photo-container{
    width: 100%;
    height: 80%;
    display: flex;
    flex-direction: column;
    align-items: center;
    overflow: scroll;
}

.photo{
    margin-bottom: 15px;
}

.footer {
    width: 100%;
    height: 10%;
    position: absolute;
    bottom: 0;
    display: flex;
    justify-content: right;
}

.footer button{
    width: 33.33%;
    color: grey;
    border: 1px solid #ccc;
}

</style>