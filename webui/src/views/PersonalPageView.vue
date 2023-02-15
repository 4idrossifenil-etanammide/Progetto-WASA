<script>
export default {
    data() {
        return {
            profileName : "",
            profileData : [],
            followerNumber : 0,
            showUploadPhoto: false,
            showChangeName: false,
            userName: ""
        }
    },
    async mounted() {
        try{
            this.userName = localStorage.getItem("name")
            this.profileName = localStorage.getItem("profile")
            let response = await this.$axios.get("/profiles/" + this.profileName)
            this.profileData = response.data
            this.followerNumber = this.profileData["Follower"] == null ? 0 : this.profileData["Follower"].length
        }catch(e){
            console.log(e)
        }
    },
    methods : {
        uploadPhoto() {
            this.showUploadPhoto = true;
        },
        closeUploadPhotoModal(bool) {
            this.showUploadPhoto = false;
            if(bool) {
                this.$router.go(0);
            }
        },
        closeUserModal(bool) {
            this.showChangeName = false
            if(bool) {
                this.$router.go(0);
            }
        },
        changeName() {
            this.showChangeName = true
        }
    }
}
</script>

<template>
    <div class="page">
        <WASAPhotoBanner></WASAPhotoBanner>
        <div class="profile-info">
            <div class="profile-name"> {{ this.profileName }} </div>
            <div class="personal-profile-component">
                <button class="change-name" @click="changeName"> Change Username </button>
                <button class="upload" @click="uploadPhoto"> Upload Photo </button>
                <div v-show="showUploadPhoto">
                    <UploadPhotoModal
                      @closePhotoModal="closeUploadPhotoModal">
                    </UploadPhotoModal>
                </div>
                <div v-show="showChangeName">
                    <ChangeNameModal
                      @closeUsernameModal="closeUserModal">
                    </ChangeNameModal>
                </div>
            </div>
            <div class="profile-info-box">
                <div class="profile-follower-box">
                    <div class="profile-follower"> Follower </div>
                    <div class="profile-follower-number"> {{ this.followerNumber }} </div>
                </div>

                <div class="profile-following-box">
                    <div class="profile-following"> Following </div>
                    <div class="profile-following-number"> {{ this.profileData["Following"] == null ? 0 : this.profileData["Following"].length }} </div>
                </div>

                <div class="profile-photoNumber-box">
                    <div class="profile-photoNumber"> Photo </div>
                    <div class="profile-photoNumber-number"> {{ this.profileData["PhotoNumber"] == null ? 0 : this.profileData["PhotoNumber"] }} </div>
                </div>
            </div>
        </div>
        <div class="photo-container">
            <div class="photo" v-show="this.profileData['Photos'] != null" v-for="photo in this.profileData['Photos']" :key="photo.photoID">
                <Photo :id="photo.photoID" :name="this.userName" :likes="photo.likeNumber" :comments="photo.commentNumber" :date="photo.date" :textComments="photo.comments"/>
            </div>
        </div>
    </div>
</template>

<style>

.personal-profile-component {
    position: relative;
    bottom: 45px;
    left: 500px;
}

.upload {
    width: 150px;
    height: 30px;
    background-color: rgb(4, 36, 177);
	color: white;
	border: none;
	border-radius: 4px;
    margin-left: 20px;
}

.change-name {
    width: 150px;
    height: 30px;
    background-color: rgb(224, 151, 15);
	color: white;
	border: none;
	border-radius: 4px;
}
</style>