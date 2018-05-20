new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        email: null, // Email address used for grabbing an avatar
        username: null, // Our username
        joined: false // True if email and username have been filled in
    },

    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
            self.chatContent +=
                '<li class="left clearfix single-message">'+
                '<span class="chat-img float-left">'+
                '<img src="'+self.gravatarURL(msg.email)+'" alt="Avatar" class="border border-3 border-dark rounded-circle mr-3"/>'+
                '</span>'+
                '<div class="chat-body clearfix">'+
                '<div class="header">'+
                '<strong class="primary-font">'+msg.username+'</strong>'+
                '</div>'+
                '<p>'+emojione.toImage(msg.message)+'</p>'+
                '</div>'+
                '</li>';

           
        });
    },

    methods: {
        send: function () {
            if (this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                        email: this.email,
                        username: this.username,
                        message: $('<p>').html(this.newMsg).text() // Strip out html
                    }
                ));
                this.newMsg = ''; // Reset newMsg
            }
        },

        join: function () {
            if (!this.email) {
                this.showAlert($("#email-alert"));
                return
            }
            if (!this.username) {
                this.showAlert($("#username-alert"));
                return
            }
            this.email = $('<p>').html(this.email).text();
            this.username = $('<p>').html(this.username).text();
            this.joined = true;
        },

        gravatarURL: function(email) {
            return 'http://www.gravatar.com/avatar/' + CryptoJS.MD5(email);
        },

        showAlert: function(target){
            target.fadeTo(50, 1).delay(2200).fadeTo(500, 0);
        }
    }
});


$(document).ready (function(){
    $("#email-alert").hide();
    $("#username-alert").hide();    
});

$("#chat-messages").bind("DOMSubtreeModified", function (){
    var element = document.getElementById('chat-messages');
    element.scrollTop = element.scrollHeight + $('.single-message').last().prop('scrollHeight'); // Auto scroll to the bottom
});