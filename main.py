def order_coffee(size, *toppings, **customizations):
    print(f"Ordering a {size} coffee.")
    
    # Handle optional positional arguments (*args)
    if toppings:
        print("Toppings added:")
        for topping in toppings:
            print(f" - {topping}")
            
    # Handle optional keyword arguments (**kwargs)
    if customizations:
        print("Custom requests:")
        for key, value in customizations.items():
            print(f" - {key}: {value}")

# Example Call
order_coffee("Large", "chocolate",  milk="oat", "slkjfal;s")
