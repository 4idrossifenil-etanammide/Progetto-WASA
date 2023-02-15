<script>
export default {
    data() {
        return {

        }
    },
    methods : {
        async uploadImage(event) {
            try{
                const file = event.target.files[0];
                await this.$axios.post("/profiles/" + localStorage.getItem("name") + "/photos", file)
                this.$emit('closePhotoModal', true)
            } catch(e) {
                console.log(e)
            }
        }
    }
}
</script>

<template>
    <div class="modal-mask">
        <div class="photo-modal-wrapper">
            <button class="close-button-photo-modal" @click="this.$emit('closePhotoModal', false)"> 
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
            <div class="modal-body">
                <input class="input-photo" type="file" ref="fileInput" @change="uploadImage" accept=".jpg, .png"/>
            </div>
        </div>
    </div>
        
</template>

<style>

.photo-modal-wrapper {
    max-width: 700px;
    width: 100%;
    height: 100%;
    max-height: 150px;
    background-color: #fff;
    box-shadow: 1px 1px 50px rgba(0,0,0, .5);
    padding: 20px;
    border-radius: 10px;
}

.modal-body {
    position: relative;
    top: 20px;
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
}

.close-button-photo-modal {
    background-color: #ffffff;
    border:  #ffff;
    position: relative;
    left: 620px;
}

</style>
