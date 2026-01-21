package ui

// Animation helpers

// Animate adds an animation class.
func (e Element) Animate(name string) Element {
	if e.Class == "" {
		e.Class = "animate-" + name
	} else {
		e.Class += " animate-" + name
	}
	return e
}

// Hover adds a hover effect class.
func (e Element) Hover(effect string) Element {
	if e.Class == "" {
		e.Class = "transition hover-" + effect
	} else {
		e.Class += " transition hover-" + effect
	}
	return e
}

// AnimationStyles contains CSS animations.
const AnimationStyles = `
@keyframes fadeIn{from{opacity:0}to{opacity:1}}
@keyframes fadeOut{from{opacity:1}to{opacity:0}}
@keyframes slideInUp{from{transform:translateY(20px);opacity:0}to{transform:translateY(0);opacity:1}}
@keyframes slideInDown{from{transform:translateY(-20px);opacity:0}to{transform:translateY(0);opacity:1}}
@keyframes scaleIn{from{transform:scale(0.9);opacity:0}to{transform:scale(1);opacity:1}}
@keyframes bounce{0%,100%{transform:translateY(0)}50%{transform:translateY(-10px)}}
@keyframes pulse{0%,100%{opacity:1}50%{opacity:0.5}}
@keyframes shake{0%,100%{transform:translateX(0)}25%{transform:translateX(-5px)}75%{transform:translateX(5px)}}
@keyframes spin{to{transform:rotate(360deg)}}

.animate-fadeIn{animation:fadeIn .3s ease-out}
.animate-fadeOut{animation:fadeOut .3s ease-out}
.animate-slideUp{animation:slideInUp .3s ease-out}
.animate-slideDown{animation:slideInDown .3s ease-out}
.animate-scaleIn{animation:scaleIn .2s ease-out}
.animate-bounce{animation:bounce 1s infinite}
.animate-pulse{animation:pulse 2s infinite}
.animate-shake{animation:shake .5s ease-in-out}
.animate-spin{animation:spin 1s linear infinite}

.transition{transition:all .2s ease}
.hover-scale:hover{transform:scale(1.05)}
.hover-lift:hover{transform:translateY(-2px);box-shadow:0 4px 12px rgba(0,0,0,.15)}
.hover-glow:hover{box-shadow:0 0 20px rgba(59,130,246,.4)}
.hover-fade:hover{opacity:.8}
`
