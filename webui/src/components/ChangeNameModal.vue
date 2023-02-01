<script>
export default {
    data() {
        return {
            usernameToSend: '',
            changeUsernameError: false,
            usernameAlreadyUsed: false
        }
    },
    methods: {
        async sendUsername() {
            this.usernameAlreadyUsed = false;
            if(this.usernameToSend.length < 3 || this.usernameToSend.length > 16) {
                this.startShaking();
            } else {
                try{
                    await this.$axios.put("/profiles/" + localStorage.getItem("name") + "/name", {
                        name : this.usernameToSend  
                    })
                    localStorage.setItem("name", this.usernameToSend)
                    localStorage.setItem("profile", this.usernameToSend)
                    this.$emit('closeUsernameModal', true)
                } catch(e) {
                    this.usernameAlreadyUsed = true;
                }
            }
        },
        startShaking() {
            this.changeUsernameError = true;
            setTimeout(() => {
                this.changeUsernameError = false;
            }, 500);
        }
    }
}
</script>

<template>
    <div class="modal-mask">
        <div class="name-modal-wrapper">
            <button class="close-button-photo-modal" @click="this.$emit('closeUsernameModal', true)"> 
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
            <div class="username-footer">
                <div class="change-username-error-label">
                    <h6 class="change-username-error-text" :class="{ shake: this.changeUsernameError }"> The username must be at least 3 characters and a maximum of 16 </h6>
                </div>
                <form class="username-form" @submit.prevent="sendUsername">
                    <input v-model="usernameToSend" class="username-form" id="commentForm" type="text" ref="username" placeholder="Insert username..." required>
                    <button class="username-button">Send</button>
                </form>
                <div class="username-already-used-label" v-show="this.usernameAlreadyUsed">
                    <h6 class="username-already-used-text"> This username belongs to another person </h6>
                </div>
            </div>
        </div>
    </div>
</template>

<style>

.username-already-used-label {
    margin-top: 45px;
    margin-left: 50px;
}

.username-already-used-text {
    color: red;
}

@keyframes shake {
  0% { transform: translate(1px, 1px) rotate(0deg); }
  10% { transform: translate(-1px, -2px) rotate(-1deg); }
  20% { transform: translate(-3px, 0px) rotate(1deg); }
  30% { transform: translate(3px, 2px) rotate(0deg); }
  40% { transform: translate(1px, -1px) rotate(1deg); }
  50% { transform: translate(-1px, 2px) rotate(-1deg); }
  60% { transform: translate(-3px, 1px) rotate(0deg); }
  70% { transform: translate(3px, 1px) rotate(-1deg); }
  80% { transform: translate(-1px, -1px) rotate(1deg); }
  90% { transform: translate(1px, 2px) rotate(0deg); }
  100% { transform: translate(1px, -2px) rotate(-1deg); }
}

.shake {
    animation: shake 0.5s linear;
}

.change-username-error-label {
    position: relative;
    right: 40px;
}

.username-button {
    position: relative;
    left: 195px;
    margin-top: 10px;
    background-color: rgb(151, 149, 32);
	color: white;
	border: none;
	border-radius: 4px;
}

.username-form {
    width: 450px;
    height: 30px;
}

.username-footer {
    width: 100%;
    height: 10px;
    position: relative;
    left: 95px;
    top: 0px;
}

.name-modal-wrapper {
    max-width: 700px;
    width: 100%;
    height: 100%;
    max-height: 190px;
    background-color: #fff;
    box-shadow: 1px 1px 50px rgba(0,0,0, .5);
    padding: 20px;
    border-radius: 10px;
}

</style>
