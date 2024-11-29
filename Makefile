.PHONY: setup
setup:
	@if [[ ! -f ./setup.sh ]]; then \
		echo "Setup is already complete. You can delete this setup make target."; \
	else \
		chmod 755 ./setup.sh && ./setup.sh && rm ./setup.sh; \
	fi
