make:
	wails build
dev:
	cd frontend & npm run build & cd .. & wails dev
