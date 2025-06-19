var socket = new WebSocket("ws://localhost:17680/_myws/_conn/");

socket.onopen = (event)=>{
	socket.send("myws,connected!");
	console.log("websocket connected.");
};

const ping = ()=>{
	socket.send("ping-pong");
};

setInterval(ping, 20000);

const board = document.getElementById("test_board");

socket.onmessage = (event)=>{
	console.log(event.data);
	if (event.data==="ping-pong"){
		return;
	}
	try{
		let cmd = JSON.parse(event.data);
		if (cmd.typ==="call"){
			if (cmd.action==="eval") {
				let ret = eval(cmd.data);
				let res = JSON.stringify( {id: cmd.id, typ: "result", data:ret} );
				socket.send(res);
			}
			return;
		}
		else if(cmd.typ === "notify") {
			if( cmd.action === "show" ){
				board.innerHTML = board.innerHTML + "<br />" + cmd.data
			}
		}
	}
	catch(err)
	{
		console.log(err)
	}
};

socket.onclose = (event)=>{
	clearInterval(ping);
	console.log("webwocket closed.");
};

socket.onerror = (event)=>{
	clearInterval(ping);
	console.log("websocket error");
	console.log(event);
}

document.getElementById("test_button").addEventListener(
	"click",
	(event) => {
		fetch("/api/calc");
	}
);

document.getElementById("test_quit").addEventListener(
	"click",
	(event) => {
		fetch("/api/quit");
	}
);

document.getElementById("open_file").addEventListener(
	"click",
	(event) => {
		fetch("/api/open_file");
	}
);
document.getElementById("save_file").addEventListener(
	"click",
	(event) => {
		fetch("/api/save_file");
	}
);
document.getElementById("open_folder").addEventListener(
	"click",
	(event) => {
		fetch("/api/open_folder");
	}
);