package ui

// CSS holds global stylesheet
var globalCSS string

// UseTailwind enables Tailwind CSS
var UseTailwind bool

// EnableTailwind turns on Tailwind CSS (pure CSS, works for dev and prod)
func EnableTailwind() {
	UseTailwind = true
	AddCSS(TailwindMinCSS)
}

// TailwindMinCSS is a minimal Tailwind subset for production
// For full Tailwind, run: npx tailwindcss -o tailwind.css --minify
const TailwindMinCSS = `
*,::after,::before{box-sizing:border-box;border:0 solid #e5e7eb}
html{line-height:1.5;-webkit-text-size-adjust:100%;font-family:ui-sans-serif,system-ui,sans-serif}
body{margin:0;line-height:inherit}
h1,h2,h3,h4,h5,h6{font-size:inherit;font-weight:inherit}
a{color:inherit;text-decoration:inherit}
b,strong{font-weight:bolder}
button,input,select,textarea{font-family:inherit;font-size:100%;margin:0;padding:0;color:inherit}
button,[type='button']{-webkit-appearance:button;background-color:transparent;background-image:none}
button,select{text-transform:none}
::-webkit-inner-spin-button,::-webkit-outer-spin-button{height:auto}
::-webkit-search-decoration{-webkit-appearance:none}
::-webkit-file-upload-button{-webkit-appearance:button;font:inherit}
img,video{max-width:100%;height:auto}
[hidden]{display:none}
.container{width:100%;margin-left:auto;margin-right:auto;padding-left:1rem;padding-right:1rem}
.sr-only{position:absolute;width:1px;height:1px;padding:0;margin:-1px;overflow:hidden;clip:rect(0,0,0,0);white-space:nowrap;border-width:0}
.fixed{position:fixed}.absolute{position:absolute}.relative{position:relative}.sticky{position:sticky}
.inset-0{inset:0}.top-0{top:0}.right-0{right:0}.bottom-0{bottom:0}.left-0{left:0}
.z-10{z-index:10}.z-20{z-index:20}.z-50{z-index:50}
.m-0{margin:0}.m-1{margin:.25rem}.m-2{margin:.5rem}.m-3{margin:.75rem}.m-4{margin:1rem}.m-6{margin:1.5rem}.m-8{margin:2rem}.m-auto{margin:auto}
.mx-auto{margin-left:auto;margin-right:auto}.my-2{margin-top:.5rem;margin-bottom:.5rem}.my-4{margin-top:1rem;margin-bottom:1rem}
.mb-1{margin-bottom:.25rem}.mb-2{margin-bottom:.5rem}.mb-3{margin-bottom:.75rem}.mb-4{margin-bottom:1rem}.mb-6{margin-bottom:1.5rem}.mb-8{margin-bottom:2rem}
.mt-1{margin-top:.25rem}.mt-2{margin-top:.5rem}.mt-4{margin-top:1rem}.mt-6{margin-top:1.5rem}.mt-8{margin-top:2rem}
.ml-2{margin-left:.5rem}.ml-4{margin-left:1rem}.mr-2{margin-right:.5rem}.mr-4{margin-right:1rem}
.block{display:block}.inline-block{display:inline-block}.inline{display:inline}.flex{display:flex}.inline-flex{display:inline-flex}.grid{display:grid}.hidden{display:none}
.h-4{height:1rem}.h-6{height:1.5rem}.h-8{height:2rem}.h-10{height:2.5rem}.h-12{height:3rem}.h-16{height:4rem}.h-full{height:100%}.h-screen{height:100vh}
.w-4{width:1rem}.w-6{width:1.5rem}.w-8{width:2rem}.w-10{width:2.5rem}.w-12{width:3rem}.w-16{width:4rem}.w-full{width:100%}.w-screen{width:100vw}
.min-h-screen{min-height:100vh}.max-w-sm{max-width:24rem}.max-w-md{max-width:28rem}.max-w-lg{max-width:32rem}.max-w-xl{max-width:36rem}.max-w-2xl{max-width:42rem}.max-w-4xl{max-width:56rem}.max-w-6xl{max-width:72rem}
.flex-1{flex:1 1 0%}.flex-col{flex-direction:column}.flex-row{flex-direction:row}.flex-wrap{flex-wrap:wrap}
.items-start{align-items:flex-start}.items-center{align-items:center}.items-end{align-items:flex-end}
.justify-start{justify-content:flex-start}.justify-center{justify-content:center}.justify-end{justify-content:flex-end}.justify-between{justify-content:space-between}
.gap-1{gap:.25rem}.gap-2{gap:.5rem}.gap-3{gap:.75rem}.gap-4{gap:1rem}.gap-6{gap:1.5rem}.gap-8{gap:2rem}
.space-x-2>:not([hidden])~:not([hidden]){margin-left:.5rem}.space-x-4>:not([hidden])~:not([hidden]){margin-left:1rem}
.space-y-2>:not([hidden])~:not([hidden]){margin-top:.5rem}.space-y-4>:not([hidden])~:not([hidden]){margin-top:1rem}
.overflow-hidden{overflow:hidden}.overflow-auto{overflow:auto}.overflow-scroll{overflow:scroll}
.rounded{border-radius:.25rem}.rounded-md{border-radius:.375rem}.rounded-lg{border-radius:.5rem}.rounded-xl{border-radius:.75rem}.rounded-2xl{border-radius:1rem}.rounded-full{border-radius:9999px}
.rounded-l{border-top-left-radius:.25rem;border-bottom-left-radius:.25rem}.rounded-r{border-top-right-radius:.25rem;border-bottom-right-radius:.25rem}
.border{border-width:1px}.border-2{border-width:2px}.border-t{border-top-width:1px}.border-b{border-bottom-width:1px}.border-l-4{border-left-width:4px}
.border-gray-200{border-color:#e5e7eb}.border-gray-300{border-color:#d1d5db}.border-blue-500{border-color:#3b82f6}.border-red-500{border-color:#ef4444}.border-green-500{border-color:#22c55e}.border-yellow-500{border-color:#eab308}
.bg-white{background-color:#fff}.bg-black{background-color:#000}.bg-transparent{background-color:transparent}
.bg-gray-50{background-color:#f9fafb}.bg-gray-100{background-color:#f3f4f6}.bg-gray-200{background-color:#e5e7eb}.bg-gray-800{background-color:#1f2937}.bg-gray-900{background-color:#111827}
.bg-red-50{background-color:#fef2f2}.bg-red-100{background-color:#fee2e2}.bg-red-500{background-color:#ef4444}.bg-red-600{background-color:#dc2626}
.bg-green-50{background-color:#f0fdf4}.bg-green-100{background-color:#dcfce7}.bg-green-500{background-color:#22c55e}.bg-green-600{background-color:#16a34a}
.bg-blue-50{background-color:#eff6ff}.bg-blue-100{background-color:#dbeafe}.bg-blue-500{background-color:#3b82f6}.bg-blue-600{background-color:#2563eb}
.bg-yellow-50{background-color:#fefce8}.bg-yellow-100{background-color:#fef9c3}.bg-yellow-500{background-color:#eab308}
.bg-purple-500{background-color:#a855f7}.bg-purple-600{background-color:#9333ea}
.bg-gradient-to-r{background-image:linear-gradient(to right,var(--tw-gradient-stops))}
.from-blue-500{--tw-gradient-from:#3b82f6;--tw-gradient-stops:var(--tw-gradient-from),var(--tw-gradient-to,rgba(59,130,246,0))}
.from-blue-600{--tw-gradient-from:#2563eb;--tw-gradient-stops:var(--tw-gradient-from),var(--tw-gradient-to,rgba(37,99,235,0))}
.to-purple-500{--tw-gradient-to:#a855f7}.to-purple-600{--tw-gradient-to:#9333ea}
.p-1{padding:.25rem}.p-2{padding:.5rem}.p-3{padding:.75rem}.p-4{padding:1rem}.p-6{padding:1.5rem}.p-8{padding:2rem}
.px-2{padding-left:.5rem;padding-right:.5rem}.px-3{padding-left:.75rem;padding-right:.75rem}.px-4{padding-left:1rem;padding-right:1rem}.px-6{padding-left:1.5rem;padding-right:1.5rem}.px-8{padding-left:2rem;padding-right:2rem}
.py-1{padding-top:.25rem;padding-bottom:.25rem}.py-2{padding-top:.5rem;padding-bottom:.5rem}.py-3{padding-top:.75rem;padding-bottom:.75rem}.py-4{padding-top:1rem;padding-bottom:1rem}
.text-left{text-align:left}.text-center{text-align:center}.text-right{text-align:right}
.text-xs{font-size:.75rem}.text-sm{font-size:.875rem}.text-base{font-size:1rem}.text-lg{font-size:1.125rem}.text-xl{font-size:1.25rem}.text-2xl{font-size:1.5rem}.text-3xl{font-size:1.875rem}.text-4xl{font-size:2.25rem}
.font-medium{font-weight:500}.font-semibold{font-weight:600}.font-bold{font-weight:700}
.leading-tight{line-height:1.25}.leading-normal{line-height:1.5}.leading-relaxed{line-height:1.625}
.text-white{color:#fff}.text-black{color:#000}
.text-gray-400{color:#9ca3af}.text-gray-500{color:#6b7280}.text-gray-600{color:#4b5563}.text-gray-700{color:#374151}.text-gray-800{color:#1f2937}.text-gray-900{color:#111827}
.text-red-500{color:#ef4444}.text-red-600{color:#dc2626}.text-red-700{color:#b91c1c}
.text-green-500{color:#22c55e}.text-green-600{color:#16a34a}.text-green-700{color:#15803d}
.text-blue-500{color:#3b82f6}.text-blue-600{color:#2563eb}.text-blue-700{color:#1d4ed8}.text-blue-200{color:#bfdbfe}
.text-yellow-600{color:#ca8a04}.text-yellow-700{color:#a16207}
.underline{text-decoration:underline}.no-underline{text-decoration:none}.line-through{text-decoration:line-through}
.opacity-0{opacity:0}.opacity-50{opacity:.5}.opacity-75{opacity:.75}.opacity-100{opacity:1}
.shadow{box-shadow:0 1px 3px 0 rgba(0,0,0,.1),0 1px 2px -1px rgba(0,0,0,.1)}.shadow-md{box-shadow:0 4px 6px -1px rgba(0,0,0,.1),0 2px 4px -2px rgba(0,0,0,.1)}.shadow-lg{box-shadow:0 10px 15px -3px rgba(0,0,0,.1),0 4px 6px -4px rgba(0,0,0,.1)}.shadow-xl{box-shadow:0 20px 25px -5px rgba(0,0,0,.1),0 8px 10px -6px rgba(0,0,0,.1)}
.outline-none{outline:2px solid transparent;outline-offset:2px}
.ring-2{box-shadow:0 0 0 2px var(--tw-ring-color)}.ring-blue-500{--tw-ring-color:#3b82f6}
.transition{transition-property:color,background-color,border-color,fill,stroke,opacity,box-shadow,transform;transition-timing-function:cubic-bezier(.4,0,.2,1);transition-duration:150ms}
.duration-200{transition-duration:200ms}.duration-300{transition-duration:300ms}
.ease-in-out{transition-timing-function:cubic-bezier(.4,0,.2,1)}
.transform{transform:var(--tw-transform)}.scale-95{--tw-scale-x:.95;--tw-scale-y:.95;transform:scale(.95)}.scale-100{--tw-scale-x:1;--tw-scale-y:1;transform:scale(1)}
.hover\:bg-gray-100:hover{background-color:#f3f4f6}.hover\:bg-gray-200:hover{background-color:#e5e7eb}
.hover\:bg-red-600:hover{background-color:#dc2626}.hover\:bg-green-600:hover{background-color:#16a34a}.hover\:bg-blue-600:hover{background-color:#2563eb}.hover\:bg-blue-700:hover{background-color:#1d4ed8}
.hover\:bg-blue-50:hover{background-color:#eff6ff}
.hover\:text-gray-900:hover{color:#111827}.hover\:text-blue-600:hover{color:#2563eb}
.hover\:shadow-lg:hover{box-shadow:0 10px 15px -3px rgba(0,0,0,.1),0 4px 6px -4px rgba(0,0,0,.1)}.hover\:shadow-xl:hover{box-shadow:0 20px 25px -5px rgba(0,0,0,.1),0 8px 10px -6px rgba(0,0,0,.1)}
.focus\:outline-none:focus{outline:2px solid transparent;outline-offset:2px}.focus\:ring-2:focus{box-shadow:0 0 0 2px var(--tw-ring-color)}.focus\:ring-blue-500:focus{--tw-ring-color:#3b82f6}
.cursor-pointer{cursor:pointer}.cursor-not-allowed{cursor:not-allowed}
.select-none{user-select:none}
.disabled\:opacity-50:disabled{opacity:.5}.disabled\:cursor-not-allowed:disabled{cursor:not-allowed}
@media(min-width:640px){.sm\:px-6{padding-left:1.5rem;padding-right:1.5rem}}
@media(min-width:768px){.md\:flex{display:flex}.md\:grid-cols-2{grid-template-columns:repeat(2,minmax(0,1fr))}.md\:grid-cols-3{grid-template-columns:repeat(3,minmax(0,1fr))}}
@media(min-width:1024px){.lg\:px-8{padding-left:2rem;padding-right:2rem}.lg\:grid-cols-3{grid-template-columns:repeat(3,minmax(0,1fr))}.lg\:grid-cols-4{grid-template-columns:repeat(4,minmax(0,1fr))}}
.grid-cols-1{grid-template-columns:repeat(1,minmax(0,1fr))}.grid-cols-2{grid-template-columns:repeat(2,minmax(0,1fr))}.grid-cols-3{grid-template-columns:repeat(3,minmax(0,1fr))}
`

// AddCSS appends to global stylesheet
func AddCSS(css string) {
	globalCSS += css + "\n"
}

// GetCSS returns the global stylesheet
func GetCSS() string {
	return globalCSS
}

// ResetCSS clears the global stylesheet
func ResetCSS() {
	globalCSS = ""
}

// Common CSS presets
const ResetStyles = `*,*::before,*::after{box-sizing:border-box;margin:0;padding:0}
body{font-family:-apple-system,BlinkMacSystemFont,sans-serif;line-height:1.5}
button,input{font:inherit}`

const BaseStyles = `
.container{max-width:800px;margin:0 auto;padding:20px}
.flex{display:flex}.flex-col{flex-direction:column}.flex-row{flex-direction:row}
.items-center{align-items:center}.justify-center{justify-content:center}.justify-between{justify-content:space-between}
.gap-1{gap:4px}.gap-2{gap:8px}.gap-3{gap:12px}.gap-4{gap:16px}
.p-1{padding:4px}.p-2{padding:8px}.p-3{padding:12px}.p-4{padding:16px}
.m-1{margin:4px}.m-2{margin:8px}.m-3{margin:12px}.m-4{margin:16px}
.text-sm{font-size:14px}.text-lg{font-size:18px}.text-xl{font-size:24px}
.font-bold{font-weight:bold}
.rounded{border-radius:4px}.rounded-lg{border-radius:8px}.rounded-full{border-radius:9999px}
.shadow{box-shadow:0 1px 3px rgba(0,0,0,0.1)}.shadow-lg{box-shadow:0 4px 6px rgba(0,0,0,0.1)}
.w-full{width:100%}.h-full{height:100%}
.cursor-pointer{cursor:pointer}
.hidden{display:none}
`

const ComponentStyles = `
.card{background:#fff;border-radius:8px;padding:16px;box-shadow:0 2px 4px rgba(0,0,0,0.1)}
.badge{display:inline-block;padding:2px 8px;border-radius:12px;font-size:12px;background:#e0e0e0}
.alert{padding:12px 16px;border-radius:6px;margin:8px 0}
.alert-info{background:#e3f2fd;color:#1565c0}
.alert-success{background:#e8f5e9;color:#2e7d32}
.alert-warning{background:#fff3e0;color:#ef6c00}
.alert-error{background:#ffebee;color:#c62828}
.modal{position:fixed;inset:0;z-index:100}
.modal-backdrop{position:fixed;inset:0;background:rgba(0,0,0,0.5);display:flex;align-items:center;justify-content:center}
.modal-content{background:#fff;border-radius:8px;padding:24px;min-width:300px;max-width:90vw;max-height:90vh;overflow:auto}
.table{width:100%;border-collapse:collapse}
.table th,.table td{padding:12px;text-align:left;border-bottom:1px solid #e0e0e0}
.table th{font-weight:600;background:#f5f5f5}
.tabs{display:flex;flex-direction:column}
.tab-list{display:flex;border-bottom:2px solid #e0e0e0}
.tab-btn{padding:8px 16px;border:none;background:none;cursor:pointer;border-bottom:2px solid transparent;margin-bottom:-2px}
.tab-btn.active{border-bottom-color:#1976d2;color:#1976d2}
.tab-content{padding:16px 0}
.dropdown{position:relative;display:inline-block}
.dropdown-trigger{cursor:pointer}
.dropdown-menu{position:absolute;top:100%;left:0;background:#fff;border-radius:6px;box-shadow:0 4px 12px rgba(0,0,0,0.15);min-width:150px;z-index:50}
.dropdown-item{padding:8px 16px;cursor:pointer}
.dropdown-item:hover{background:#f5f5f5}
.progress{height:8px;background:#e0e0e0;border-radius:4px;overflow:hidden}
.progress-bar{height:100%;background:#1976d2;transition:width 0.3s}
.avatar{width:40px;height:40px;border-radius:50%;object-fit:cover}
.spinner{width:24px;height:24px;border:3px solid #e0e0e0;border-top-color:#1976d2;border-radius:50%;animation:spin 0.8s linear infinite}
@keyframes spin{to{transform:rotate(360deg)}}
.divider{border:none;border-top:1px solid #e0e0e0;margin:16px 0}
.btn{padding:8px 16px;border:none;border-radius:6px;cursor:pointer;font-weight:500}
.btn-primary{background:#1976d2;color:#fff}
.btn-secondary{background:#e0e0e0;color:#333}
.btn-danger{background:#c62828;color:#fff}
`
