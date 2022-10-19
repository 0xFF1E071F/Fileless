build_src: 
	@echo "[+] - Building the bootstrap server"
	go build -o build/bootstrap src/bootstrap/*.go
	
	@echo "[+] - Building the agent"
	go build -o build/agent src/agent/*.go
	
	@echo "[+]  - Copying the stager"
	cp stager/stager.py build/stager.py

	@echo "Building finished!"



