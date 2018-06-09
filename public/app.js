var saveCookie = function(e){
    Cookies.set('username', app.username, { expires: 1/4320 });
    Cookies.set('email', app.email, { expires: 1/4320 });
    Cookies.set('joined', app.joined, { expires: 1/4320 });
};

var isChatBottom = true;

var checkChatBottom = function(){
    var element = document.getElementById('chat-messages');
    isChatBottom = ((element.scrollTop + element.clientHeight) >= element.scrollHeight)
    return isChatBottom
}

var scrollChatToBottom = function(){
    if(!isChatBottom){
        return 
    }
    var element = document.getElementById('chat-messages');
    element.scrollTop = element.scrollHeight 
}

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
        prevMsg: '',
        msgCount: 0,
        isNotice: false,
    },

    created: function() {
        var audio = new Audio('receive.wav')
        var self = this;
        this.ws = new WebSocket('wss://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
            var msgId = 'msg-' + self.msgCount;

            if (msg.message != "#userlistjoin#") {
                var msgContent = $('<p>' + emojione.toImage(msg.message) + '</p>').linkify({target: "_blank"}); //$('li.single-message:last p').linkify({target: "_blank"});

                $(msgContent).find('a.linkified').each(function( index, element ) {
                    var imgSource = $(element).attr('href');
                    self.getMeta(imgSource, function(width, height) { 
                        if(((width || 0) == 0) || ((height || 0) == 0)){
                            return
                        }
    
                        var chatBodyHeight = $('#chat-body').height();
                        var imgClass = '';
    
                        if(height*10/19 < chatBodyHeight){
                            imgClass = 'chat-image';
                        } else if (height*10/39 < chatBodyHeight){
                            imgClass = 'chat-image-half';
                        }
    
                        if (imgClass != ''){
                            $(element).text( '' );
                            $(element).append( '<img class="' + imgClass + '" src="' + imgSource + '">' );    
                        }

                        $('#'+msgId).html(msgContent.html());
                        self.chatContent = $('#chat-messages').html();
                    });
                });

                var img = self.getImageURL(msg.email);
                checkChatBottom();
                self.chatContent +=
                    '<li class="left clearfix single-message">'+
                    '<span class="chat-img float-left">'+
                    '<img src="'+ img +'" alt="Avatar" class="border border-3 border-dark rounded-circle img-sizing"/>'+
                    '</span>'+
                    '<div class="chat-body clearfix">'+
                    '<div class="header">'+
                    '<strong class="primary-font">'+msg.username+'</strong>'+
                    '</div>'+
                    '<p id="' + msgId + '">'+msgContent.html()+'</p>'+
                    '</div>'+
                    '</li>';

                if (self.isNotice){
                    Notification.requestPermission(function(result) {
                        if (result === 'granted') {
                            navigator.serviceWorker.ready.then(function(registration) {
                                registration.showNotification(msg.username, { body: msgContent.text(), icon: img, vibrate: [100] })
                            });
                        }
                    });
                    audio.play();
                }

                self.msgCount++;
            }

            if(msg.email == "Manager@chat.com"){
                var message = msg.message;
                var username = msg.information;
                if(message.includes("join")){
                    self.chatMember += 
                        '<li class="p-2 m-0" id="' + self.slugify(username) + '">' + username + '</li>';
                }else if(message.includes("set")){
                    self.removeChatMember('unknown-user')
                    self.chatMember += 
                        '<li class="p-2 m-0" id="' + self.slugify(username) + '">' + username + '</li>';
                }else if(message.includes("left")){
                    self.removeChatMember(username)
                }else if(message.includes("change")){
                    usernames = username.split(',');
                    self.removeChatMember(usernames[0])
                    self.chatMember += 
                        '<li class="p-2 m-0" id="' + self.slugify(usernames[1]) + '">' + usernames[1] + '</li>';
                }
            }
        });
        this.ws.addEventListener('open', function(e){
            self.username = Cookies.get('username') || '';
            self.email = Cookies.get('email') || '';
            self.joined = ((Cookies.get('joined') || 'false') == 'true');
        
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
            return 'https://www.gravatar.com/avatar/' + CryptoJS.MD5(email);
        },

        showAlert: function(target){
            target.fadeTo(50, 1).delay(2200).fadeTo(500, 0);
        },

        removeChatMember: function(username){
            var html = $.parseHTML(this.chatMember);
            html = $(html).not('#'+this.slugify(username)+':first');
            var s = '';
            $(html).each(function(){
                s += $(this).clone().wrap('<ul>').parent().html();
            });
            this.chatMember = s;
        },

        getMeta: function(url, callback) {
            var img = new Image();
            img.src = url;
            img.onload = function() { callback(this.width, this.height); }
        },

        slugify: function(text) {
            var slug = text.toLowerCase().trim()
                .replace(/[\s`~!@#$%^&*()_|+\-=?;:'",.<>\{\}\[\]\\\/]+/g, '-')       // Swap any length of whitespace, underscore, hyphen characters with a single -
                .replace(/^-+|-+$/g, '');       // Remove leading, trailing -

            if( slug == '' ){
                slug = 'no-id'
            }

            return slug
        }
    }
});

var chatBottomFix = function(){
    checkChatBottom();
    scrollChatToBottom();
}

$(window).on('load', function(){
    $("#username-alert").hide();
    
    $(window).resize(chatBottomFix);
    $("#chat-messages").resize(chatBottomFix);

    $("#side-in-button").on("click", function(e){
        checkChatBottom();
        $("#side-bar").removeClass(['d-md-block', 'col-md-4', 'col-lg-3', 'col-xl-3']);
        $("#side-bar").addClass(['d-md-none']);
        $("#main-content").removeClass(['col-md-8', 'col-lg-9', 'col-xl-9']);
        $("#side-out-button").removeAttr("hidden");
        scrollChatToBottom();
    })

    $("#side-out-button").on("click", function(e){
        checkChatBottom();
        $("#side-bar").addClass(['d-md-block', 'col-md-4', 'col-lg-3', 'col-xl-3']);
        $("#side-bar").removeClass(['d-md-none']);
        $("#main-content").addClass(['col-md-8', 'col-lg-9', 'col-xl-9']);
        $("#side-out-button").attr("hidden", "hidden");
        scrollChatToBottom();
    })

    $("#top-in-button").on("click", function(e){
        $("#top-out-button").removeAttr("hidden");
    })

    $("#top-out-button").on("click", function(e){
        $("#top-out-button").attr("hidden", "hidden");
    })

    $("#chat-messages").bind("DOMSubtreeModified", function (){
        if (isChatBottom){
            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight + $('.single-message').last().prop('scrollHeight'); // Auto scroll to the bottom
        }
    });

    Notification.requestPermission(function (status) {
        if (Notification.permission !== status) {
            Notification.permission = status;
        }
    });  

    $(window).blur(function(){
        app.isNotice = true;
    });
    $(window).focus(function(){
        app.isNotice = false;
    });
    $(window).on('unload', saveCookie);
});

navigator.serviceWorker.register('sw.js');