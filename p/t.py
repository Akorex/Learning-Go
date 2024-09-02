import time

start_time = time.time()

for i in range(100000):
    print(i)

end_time = time.time()
time_taken = end_time - start_time

print(f"Time taken {time_taken}")