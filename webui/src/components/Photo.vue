<script>
export default {
    data(){
        return{
            photoURL: __API_URL__ + "/profiles/" + localStorage.getItem("name") + "/photos/"
        }
    },

    props: ["id", "name", "likes", "comments"]
}
</script>

<template>
    <div class="rounded-image-container">
        <img :src="photoURL + id" class="image">
        <div class="upper-left">
            {{ name }}
        </div>
        <div class="bottom-right">
            <div class="like-count"> {{ likes }} likes</div>
            <div class="comment-count"> {{ comments }} comments</div>
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
    bottom: 10px;
    right: 10px;
    font-size: 14px;
}

.like-count{
    margin-bottom: 5px;
}
</style>