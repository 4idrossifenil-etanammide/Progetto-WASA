<script>
export default{
    data() {
        return {
            photoComments: []
        }
    },
    async mounted(){
        this.photoComments = this.comments
        console.log(this.comments)
    },
    methods: {
        sendComment() {

        }
    },

    props: ["comments"]
}
</script>

<template id="modal-template">
    <transition name="modal">
        <div class="modal-mask">
            <div class="modal-wrapper">
                <div class="header">
                    <div class="title"> Commenti </div>
                    <button class="close-button" @click="this.$emit('closeCommentModal', false)"> 
                        <svg width="24" height="24" viewBox="0 0 512.000000 512.000000">
                            
                            <g transform="translate(0.000000,512.000000) scale(0.100000,-0.100000)"
                            fill="#000000" stroke="none">
                                <path d="M3735 4157 c-22 -8 -51 -19 -65 -26 -14 -7 -269 -257 -567 -554
                                l-543 -542 -532 531 c-348 346 -550 541 -583 560 -46 26 -58 29 -145 29 -84 0
                                -101 -3 -145 -27 -67 -35 -140 -110 -173 -178 -22 -46 -27 -68 -27 -135 0 -67
                                5 -88 27 -131 21 -39 160 -184 565 -587 l538 -536 -536 -538 c-614 -616 -595
                                -592 -587 -736 5 -97 33 -158 106 -232 71 -71 140 -100 237 -100 67 0 88 5
                                131 27 38 20 187 163 587 565 l536 538 543 -542 c603 -603 584 -587 714 -587
                                60 0 81 5 136 32 130 64 208 184 208 320 0 128 10 115 -588 715 l-536 537 545
                                548 c608 609 584 579 584 717 0 59 -5 81 -32 136 -59 119 -173 198 -297 205
                                -34 2 -79 -2 -101 -9z"/>
                            </g>
                        </svg>
                    </button>
                </div>
                <div class="comment-container">
                    <div class="comment-section" v-for="c in this.photoComments">
                        <div class="comment-name"> {{ c.name }} </div>
                        <div class="comment-text"> {{ c.comment }} </div>
                    </div>
                </div>
                <div class="footer">
                    <form class="comment-form" @submit.prevent="sendComment">
                        <input v-model="comment" type="text" ref="username" placeholder="Insert comment..." required>
                        <button>Send</button>
                    </form>
                </div>
            </div>
        </div>
    </transition>
</template>

<style>

.comment-section {
    margin-bottom: 10px;
}

.comment-name {
    position: relative;
    left: 5px;
    font-weight: bold;
    font-size: 15px;
}

.comment-text {
    position: relative;
    left: 15px;
    font-size: 13px;
}

.modal-mask{
    position: fixed;
    z-index: 9998;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0,0,0, .5);
    display: flex;
    justify-content: center;
    align-items: center;
}

.modal-wrapper {
    max-width: 700px;
    width: 100%;
    height: 100%;
    max-height: 500px;
    background-color: #fff;
    box-shadow: 1px 1px 50px rgba(0,0,0, .5);
    padding: 20px;
    border-radius: 10px;
}

.title{
    font-weight: bold;
    font-size: 30px;
}

.header {
    position: relative;
    display: flex;
    flex-direction: row;
}

.close-button {
    background-color: #ffffff;
    border:  #ffff;
    position: absolute;
    top: 10px;
    right: 5px;
}

.comment-container {
    width: 100%;
    height: 80%;
    margin-bottom: 10px;
    display: flex;
    flex-direction: column;
    align-items: left;
    overflow: scroll;
}

.footer {
    height: 10px;
    position: relative;
    bottom: 0;
}

.comment-form {
    width: 100%;
    height: 30px;
    display: flex;
    flex-direction: row;
}

.comment-form input {
    width: 90%;
}

.comment-form button {
    width: 10%;
}

</style>
