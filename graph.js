var graphNodes = document.getElementsByClassName('graphNode')

for(var i in graphNodes){
	if(Number.isInteger(parseInt(i))){
		var bar = new ProgressBar.Circle(graphNodes[i], {
		  color: 'whitesmoke',
		  strokeWidth: 12,
		  trailWidth: 2,
		  easing: 'easeInOut',
		  duration: 1400,
		  text: {
		    autoStyleContainer: false
		  },
		  from: { color: '#33e', width: 12 },
		  to: { color: '#3be', width: 12 },
		  step: function(state, circle) {
		  	circle.trail.setAttribute('stroke', '#333')
		  	circle.trail.setAttribute('fill-opacity', '0.8')
		    circle.path.setAttribute('stroke', state.color);
		    circle.path.setAttribute('stroke-width', state.width);

		    var value = Math.round(circle.value() * 100);
		    if (value === 0) {
		      circle.setText('');
		    } else {
		      circle.setText(Math.floor(value / 10));
		    }

		  }
		});
		bar.text.style.fontFamily = '"Raleway", Helvetica, sans-serif';
		bar.text.style.fontSize = '2.5rem';
		bar.text.style.top = '45%';
		bar.text.style.color = 'whitesmoke'

		bar.animate(graphNodes[i].dataset.val / 10);
	}
}