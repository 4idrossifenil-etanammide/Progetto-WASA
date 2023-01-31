<script>

export default {
    data: function() {
        return {
            userToSearch: ''
        }
    },
    methods: {
        search() {
            console.log(this.$route.path)
            this.userToSearch = this.$refs.userToSearch.value;
            localStorage.setItem("profile", this.userToSearch)
            if(this.$route.path == "/homepage/profile") {
                this.$emit('updateParent')
            } else {
                this.$router.replace("/homepage/profile");
            }
        },
        goPersonalProfile() {
            console.log("AAAAAA")
            localStorage.setItem("profile", localStorage.getItem("name"))
            this.$router.replace("/homepage/profile/personal");
        }
    }
}

</script>

<template>
    <div class="header">
        <h1 class="title">WasaPHOTO</h1>
        <form class="search-form" @submit.prevent="search">
			<input class="input-form" v-model="userToSearch" ref="userToSearch" type="text" placeholder="Search username..." required>
		</form>
        <button class="profile-button" @click="goPersonalProfile">
            <svg version="1.0" width="40" height="40" viewBox="0 0 612.000000 612.000000">
                
                <g transform="translate(0.000000,612.000000) scale(0.100000,-0.100000)"
                fill="#000000" stroke="none">
                <path d="M2885 5055 c-390 -71 -690 -353 -792 -743 -36 -140 -36 -334 1 -477
                114 -446 509 -755 966 -755 655 0 1132 626 959 1260 -118 433 -516 733 -968
                729 -53 0 -127 -7 -166 -14z"/>
                <path d="M2244 3084 c-313 -160 -566 -439 -698 -770 -35 -89 -185 -648 -207
                -769 -37 -213 120 -439 338 -485 32 -7 505 -10 1413 -8 l1365 3 56 23 c80 34
                146 88 196 160 79 116 95 239 54 402 -87 339 -157 593 -178 649 -122 327 -381
                622 -683 780 l-98 51 -64 -44 c-392 -273 -947 -275 -1348 -6 -30 20 -60 39
                -66 43 -6 3 -42 -10 -80 -29z"/>
                </g>
            </svg>
        </button>
    </div>
</template>

<style>

.profile-button {
    position: relative;
    right: 1050px;
    top: 5px;
    background-color: rgb(215, 255, 255);
    border:  #dfdbdb;
    margin-bottom: 7px;
}

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

</style>