var app = new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        chatMember: '<li class="p-2 m-0 font-weight-bold text-center">User list</li>',
        email: null, // Email address used for grabbing an avatar
        username: null, // Our username
        changedUsername: null,
        joined: false, // True if email and username have been filled in
        prevMsg: ''
    },

    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);


            if (msg.message != "#userlistjoin#") {
                self.chatContent +=
                    '<li class="left clearfix single-message">'+
                    '<span class="chat-img float-left">'+
                    '<img src="'+self.getImageURL(msg.email)+'" alt="Avatar" class="border border-3 border-dark rounded-circle img-sizing"/>'+
                    '</span>'+
                    '<div class="chat-body clearfix">'+
                    '<div class="header">'+
                    '<strong class="primary-font">'+msg.username+'</strong>'+
                    '</div>'+
                    '<p>'+emojione.toImage(msg.message)+'</p>'+
                    '</div>'+
                    '</li>';
            }

            if(msg.email == "Manager@chat.com"){
                var message = msg.message;
                var username = msg.information;
                if(message.includes("join") > 0){
                    self.chatMember += 
                        '<li class="p-2 m-0" id="' + username.replace(' ','-') + '">' + username + '</li>';
                }else if(message.includes("set") > 0){
                    self.removeChatMember('unknown-user')
                    self.chatMember += 
                        '<li class="p-2 m-0" id="' + username.replace(' ','-') + '">' + username + '</li>';
                }else if(message.includes("left") > 0){
                    self.removeChatMember(username)
                }else if(message.includes("change") > 0){
                    usernames = username.split(',');
                    self.removeChatMember(usernames[0])
                    self.chatMember += 
                        '<li class="p-2 m-0" id="' + usernames[1].replace(' ','-') + '">' + usernames[1] + '</li>';
                }
            }
        });
        this.ws.addEventListener('open', function(e){
            self.username = Cookies.get('username') || '';
            self.email = Cookies.get('email') || '';
            self.joined = ((Cookies.get('joined') || 'false') == 'true');
        
            console.log(self.username);
            console.log(self.email);
            console.log(self.joined);

            self.ws.send(
                JSON.stringify({
                    email: self.email == '' ? 'anonymous@mail.com' : self.email,
                    username: self.username == '' ? 'unknown user' : self.username,
                    message: '#login#',
                    information: ''
                }
            ));
        })

    },

    updated: function(){
        $('li.single-message p').linkify({
            target: "_blank"
        });
        if(this.prevMsg != '' && this.newMsg == this.prevMsg){
            this.newMsg = '';
            this.prevMsg = '';
        }
        if(this.newMsg == ''){
            this.prevMsg = '';
        }
    },

    methods: {
        send: function () {
            if (this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                        email: this.email,
                        username: this.username,
                        message: $('<p>').html(this.newMsg).text(),
                        information: ''
                    }
                ));
            }
            this.prevMsg = this.newMsg;
            this.newMsg = '';
        },

        getImageURL: function(email){
            if(email == 'Manager@chat.com'){
                return 'info.png';
            }else{
                return this.gravatarURL(email)
            }
        },

        join: function () {
            if (!this.email) {
                this.email = "anonymous@mail.com"
            }
            if (!this.username) {
                this.showAlert($("#username-alert"));
                return
            }
            this.email = $('<p>').html(this.email).text();
            this.username = $('<p>').html(this.username).text();
            this.joined = true;
            this.ws.send(
                JSON.stringify({
                    email: this.email,
                    username: this.username,
                    message: '#join#',
                    information: ''
                }
            ));
        },

        change: function(){
            if (!this.changedUsername) {
                this.showAlert($("#username-alert"));
                return
            }
            this.changedUsername = $('<p>').html(this.changedUsername).text();
            this.username = $('<p>').html(this.username).text();
            this.ws.send(
                JSON.stringify({
                    email: this.changedUsername,
                    username: this.username,
                    message: '#change#',
                    information: ''
                }
            ));
            this.username = this.changedUsername;
            this.changedUsername = '';
        },

        clear: function(){
            this.chatContent = '';
        },

        gravatarURL: function(email) {
            return 'http://www.gravatar.com/avatar/' + CryptoJS.MD5(email);
        },

        showAlert: function(target){
            target.fadeTo(50, 1).delay(2200).fadeTo(500, 0);
        },

        removeChatMember: function(username){
            var html = $.parseHTML(this.chatMember);
            html = $(html).not('#'+username.replace(' ','-')+':first');
            var s = '';
            $(html).each(function(){
                s += $(this).clone().wrap('<ul>').parent().html();
            });
            this.chatMember = s;
        }
    }
});

var saveCookie = function(e){
    Cookies.set('username', app.username, { expires: 1/4320 });
    Cookies.set('email', app.email, { expires: 1/4320 });
    Cookies.set('joined', app.joined, { expires: 1/4320 });
};

$(window).on('load', function(){
    $("#username-alert").hide();
    
    var _originalSize = $(window).width() + $(window).height()
    $(window).resize(function(){
        var element = document.getElementById('chat-messages');
        element.scrollTop = element.scrollHeight 
    });

    $("#side-in-button").on("click", function(e){
        $("#side-bar").removeClass(['d-md-block', 'col-md-4', 'col-lg-3', 'col-xl-3']);
        $("#side-bar").addClass(['d-md-none']);
        $("#main-content").removeClass(['col-md-8', 'col-lg-9', 'col-xl-9']);
        $("#side-out-button").removeAttr("hidden");
    })

    $("#side-out-button").on("click", function(e){
        $("#side-bar").addClass(['d-md-block', 'col-md-4', 'col-lg-3', 'col-xl-3']);
        $("#side-bar").removeClass(['d-md-none']);
        $("#main-content").addClass(['col-md-8', 'col-lg-9', 'col-xl-9']);
        $("#side-out-button").attr("hidden", "hidden");
    })

    $("#top-in-button").on("click", function(e){
        $("#top-out-button").removeAttr("hidden");
    })

    $("#top-out-button").on("click", function(e){
        $("#top-out-button").attr("hidden", "hidden");
    })


    $(window).on('unload', saveCookie);
});

$("#chat-messages").bind("DOMSubtreeModified", function (){
    var element = document.getElementById('chat-messages');
    element.scrollTop = element.scrollHeight + $('.single-message').last().prop('scrollHeight'); // Auto scroll to the bottom
});