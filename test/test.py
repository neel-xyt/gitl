import os
import shutil
import subprocess
import tempfile
import sys

def run_in_isolated_env(binary_path: str, args: list = []):
    binary_path = os.path.abspath(binary_path)

    if not os.path.isfile(binary_path):
        print(f"[-] Binary not found: {binary_path}")
        return

    with tempfile.TemporaryDirectory() as temp_env:
        print(f"[*] Created isolated environment: {temp_env}")

        # Copy binary into isolated env
        binary_name = os.path.basename(binary_path)
        temp_binary = os.path.join(temp_env, binary_name)
        shutil.copy(binary_path, temp_binary)
        os.chmod(temp_binary, 0o755)

        # Optional: inject test data directory
        # os.makedirs(os.path.join(temp_env, "data"), exist_ok=True)

        os.chdir(temp_env)
        print(f"[*] Running {binary_name} inside isolated env...\n")

        try:
            subprocess.run([f"./{binary_name}"] + args, check=True)
        except subprocess.CalledProcessError as e:
            print(f"[!] Error: {e}")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python run_isolated.py <binary_path> [args...]")
        sys.exit(1)

    binary = sys.argv[1]
    extra_args = sys.argv[2:]
    run_in_isolated_env(binary, extra_args)
