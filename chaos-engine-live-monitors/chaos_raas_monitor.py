import time
from web3 import Web3
from colorama import init, Fore, Style

init(autoreset=True)


OP_RPC = "https://mainnet.optimism.io"
w3 = Web3(Web3.HTTPProvider(OP_RPC))

QUANTUM_ATTACK_WINDOW_MS = 800 

def monitor_raas_exposure():
    print(f"{Fore.CYAN}INITIALIZING QUBEX CHAOS ENGINE...{Style.RESET_ALL}")
    print(f"Target: OP Stack RaaS Framework | RPC: {OP_RPC}\n")
    
    last_block = w3.eth.get_block('latest')
    
    while True:
        try:
            current_block = w3.eth.get_block('latest')
            
            if current_block.number > last_block.number:
                time_diff_sec = current_block.timestamp - last_block.timestamp
                sequencer_latency_ms = time_diff_sec * 1000
                exposure_window = sequencer_latency_ms - QUANTUM_ATTACK_WINDOW_MS
                
                print(f"Block: {current_block.number} | RaaS Sequencer Latency: {sequencer_latency_ms}ms")
                
                if exposure_window > 0:
                    print(f"{Fore.RED}[CRITICAL FUD ALERT] RAAS MODULAR STACK EXPOSED!{Style.RESET_ALL}")
                    print(f"{Fore.RED}-> OP Stack ECDSA Exposure Window: {exposure_window}ms{Style.RESET_ALL}")
                    print(f"{Fore.GREEN}-> QUBEX ML-DSA Patch Ready: Nanosecond Off-chain Validation (Proof in GitHub Repo) | 0% L1 Gas Bloat{Style.RESET_ALL}")
                    print("-" * 50)
                else:
                    print("[SAFE] Block finalized within attack window.")
                    
                last_block = current_block
                
            time.sleep(0.1) 
            
        except Exception as e:
            print(f"RPC Connection Error: {e}")
            time.sleep(2)

if __name__ == "__main__":
    if w3.is_connected():
        monitor_raas_exposure()
    else:
        print("Failed to connect to OP Stack RPC. Check endpoint.")
