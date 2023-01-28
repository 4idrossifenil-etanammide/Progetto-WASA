<script>
export default {
    data(){
        return{
            photoURL: __API_URL__ + "/profiles/" + localStorage.getItem("name") + "/photos/",
            heartColor: "black",
            nLike : 0,
            prova: {}
        }
    },
    async mounted() {
        this.nLike = this.likes;
        await this.initializePhoto();
    },
    methods: {
        putLike() {
            this.heartColor = this.heartColor == "black" ? "red" : "black";
            this.nLike = this.heartColor == "red" ? this.nLike + 1 : this.nLike - 1;
        },
        async initializePhoto() {
            try {
                let risp = await this.$axios.get("/profiles/" + this.name + "/photos/" + this.id + "/likes/" + localStorage.getItem("name"))
                this.prova = risp.data["liked"]
                this.heartColor = this.prova == "true" ? "red" : "black"
            } catch (e) {}
        }
    },

    props: ["id", "name", "likes", "comments", "date"]
}
</script>

<template>
    <div class="rounded-image-container">
        <img :src="photoURL + id" class="image">
        <div class="upper-left">
            {{ name }}
        </div>
        <div class="bottom-left">
            {{ date.split(".")[0].split("T")[0] }} {{ date.split(".")[0].split("T")[1] }}
        </div>
        <div class="bottom-right">
            <button class="like-button" @click="putLike">
                <svg :style="{ fill: this.heartColor }" width="24" height="24">
                    <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"></path>
                </svg>
            </button>
            <div class="like-count">{{ nLike }} </div>
            <button class="comment-button">comments</button>
            <div class="comment-count">{{ comments }} </div>
        </div>
    </div>
</template>

<style>

.rounded-image-container {
    width: 450px;
    height: 300px;
    border-radius: 10px;
    overflow: hidden;
    position: relative;
    background-color: #dfdbdb;
}

.image {
    width: 100%;
    height: 70%;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
}

.upper-left {
    position: absolute;
    top: 10px;
    left: 10px;
    font-size: 18px;
    font-weight: bold;
}

.bottom-right {
    position: absolute;
    display: flex;
    flex-direction: row;
    bottom: 0px;
    right: 10px;
    font-size: 14px;
}

.bottom-left {
    position: absolute;
    bottom: 10px;
    left: 10px;
    font-size: 14px;
}

.like-count{
    margin-top: 1px;
}

.comment-count {
    margin-top: 1px;
}

.like-button {
    background-color: #dfdbdb;
    border:  #dfdbdb;
    margin-bottom: 5px;
}

.comment-button {
    background-color: #dfdbdb;
    border:  #dfdbdb;
    margin-bottom: 5px;
}
</style>