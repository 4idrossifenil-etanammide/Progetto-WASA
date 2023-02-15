<script>
export default {
    data(){
        return{
            photoURL: __API_URL__ + "/profiles/" + localStorage.getItem("name") + "/photos/",
            heartColor: "black",
            nLike : 0,
            nComments : 0,
            initColor: {},
            showModal: false,
            userName: ""
        }
    },
    async mounted() {
        this.nLike = this.likes;
        this.nComments = this.comments;
        this.userName = localStorage.getItem("name")
        await this.initializePhoto();
    },
    methods: {
        async putLike() {
            try {
                if (this.heartColor == "red") {
                    this.heartColor = "black"
                    this.nLike = this.nLike - 1
                    await this.$axios.delete("/profiles/" + this.name + "/photos/" + this.id + "/likes/" + localStorage.getItem("name"))
                } else {
                    this.heartColor = "red"
                    this.nLike = this.nLike + 1
                    await this.$axios.put("/profiles/" + this.name + "/photos/" + this.id + "/likes/" + localStorage.getItem("name"), {
                        name : localStorage.getItem("name")
                    })
                }
            } catch(e) {
                console.log(e)
            }
        },
        async initializePhoto() {
            try {
                let risp = await this.$axios.get("/profiles/" + this.name + "/photos/" + this.id + "/likes/" + localStorage.getItem("name"))
                this.initColor = risp.data["liked"]
                this.heartColor = this.initColor == "true" ? "red" : "black"
            } catch (e) {
                console.log(e)
            }
        },
        closeComment(bool) {
            this.showModal = bool;
            this.$router.go(0);
        },
        updateNComment(bool) {
            if (bool) {
                this.nComments = this.nComments + 1
            } else {
                this.nComments = this.nComments - 1
            }
        },
        async deletePhoto() {
            try{
                await this.$axios.delete("/profiles/" + localStorage.getItem("name") + "/photos/" + this.id)
                this.$router.go(0);
            } catch(e) {
                console.log(e)
            }
        }
    },

    props: ["id", "name", "likes", "comments", "date", "textComments"]
}
</script>

<template>
    <div class="rounded-image-container">
        <img :src="photoURL + id" class="image">
        <div class="upper-left">
            {{ name }}
            <div v-show="name == this.userName">
                <button class="delete-photo-button" @click="deletePhoto">
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
            <button class="comment-button" @click="showModal = true">
                <svg width="24" height="24" viewBox="0 0 1300 1300">
                    
                    <g transform="translate(-450.000000,1560.000000) scale(0.100000,-0.100000)"
                    fill="#000000" stroke="none">
                    <path d="M11735 14304 c-850 -34 -1611 -227 -2335 -589 -518 -260 -1013 -636
                    -1361 -1035 -503 -577 -798 -1226 -880 -1935 -18 -154 -15 -538 5 -704 82
                    -683 380 -1331 867 -1886 93 -105 301 -314 383 -384 l38 -32 -7 -77 c-17 -197
                    -67 -468 -119 -649 -178 -617 -506 -1035 -1007 -1279 -101 -50 -133 -71 -149
                    -96 -33 -54 -35 -92 -9 -145 30 -59 74 -87 138 -85 62 1 255 42 484 103 935
                    247 1817 659 2481 1159 l88 67 82 -23 c119 -34 379 -90 551 -118 347 -57 596
                    -77 965 -77 699 1 1344 114 1992 351 831 303 1576 832 2064 1465 415 538 660
                    1132 724 1759 16 156 14 506 -4 666 -71 625 -305 1187 -725 1735 -267 349
                    -596 649 -1031 941 -688 462 -1603 772 -2510 849 -168 14 -580 25 -725 19z
                    m662 -320 c821 -72 1554 -296 2205 -675 513 -298 1009 -754 1305 -1199 282
                    -424 457 -885 515 -1355 19 -159 16 -566 -6 -722 -90 -650 -385 -1251 -864
                    -1761 -704 -751 -1699 -1240 -2852 -1401 -80 -11 -212 -26 -295 -33 -192 -15
                    -749 -15 -925 1 -324 29 -686 93 -992 175 -140 37 -173 43 -202 35 -19 -5 -99
                    -57 -178 -116 -576 -427 -1158 -726 -1903 -979 -71 -24 -131 -44 -134 -44 -3
                    0 11 17 30 38 173 182 361 511 466 817 96 282 168 647 189 971 5 82 3 106 -10
                    131 -8 17 -75 85 -148 151 -338 305 -555 571 -758 928 -355 625 -475 1357
                    -334 2044 178 864 725 1623 1584 2197 581 388 1349 666 2121 768 355 46 850
                    59 1186 29z"/>
                    </g>
                </svg>
            </button>
            <div v-show="showModal" >
                <CommentModal 
                    @closeCommentModal="closeComment" 
                    @updateCommentCounter="updateNComment"
                    :comments="textComments" 
                    :user="name" 
                    :photoID="id">
                </CommentModal>
            </div>
            <div class="comment-count">{{ this.nComments }} </div>
        </div>
    </div>
</template>

<style>

.delete-photo-button {
    position: relative;
    bottom: 30px;
    left: 400px;
    background-color: #dfdbdb;
    border:  #dfdbdb;
}

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
    margin-bottom: 7px;
}

.comment-button {
    background-color: #dfdbdb;
    border:  #dfdbdb;
    margin-bottom: 5px;
}
</style>