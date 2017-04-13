import tensorflow as tf
import random as rn
import pandas as pd
from sklearn.cross_validation import train_test_split
import numpy as np     


     
  
  

# read csv file
MNIST = pd.read_csv("C:/Users/Greg/Anaconda3/envs/tensorflow/Scripts/dummydata.csv")

# pop label column and create training label array
train_label = MNIST.pop("label")

# converts from dataframe to np array
MNIST=MNIST.values

# convert train labels to one hots
train_labels = pd.get_dummies(train_label)
# make np array
train_labels = train_labels.values

x_train,x_test,y_train,y_test = train_test_split(MNIST,train_labels,test_size=0.2)
# we now have features (x_train) and y values, separated into test and train

# convert to dtype float 32
x_train,x_test,y_train,y_test = np.array(x_train,dtype='float32'), np.array(x_test,dtype='float32'),np.array(y_train,dtype='float32'),np.array(y_test,dtype='float32')

  
  
  
epochs_completed = 0
index_in_epoch = 0
num_examples = x_train.shape[0]
num_columns = x_train.shape[1]
  
  
  
  
  
# Parameters
learning_rate = 0.001
training_iters = 20
batch_size = 10
display_step = 10

# Network Parameters
n_input = 10 #  dta input (img shape: 28*28)
n_classes = 2 # total classes (0-1 digits)
dropout = 0.75 # Dropout, probability to keep units

# tf Graph input
x = tf.placeholder(tf.float32, [n_input,num_columns])
y = tf.placeholder(tf.int16, [n_input, n_classes])
keep_prob = tf.placeholder(tf.float32) #dropout (keep probability)



def next_batch(batch_size):
    global x_train
    global y_train
    global index_in_epoch
    global epochs_completed
    start = index_in_epoch
    index_in_epoch += batch_size
    # when all trainig data have been already used, it is reorder randomly     
    if index_in_epoch > num_examples:
        # finished epoch
        epochs_completed += 1
        # shuffle the data
        perm = np.arange(num_examples)
        np.random.shuffle(perm)
        x_train = x_train[perm]
        y_train = y_train[perm]
        # start next epoch
        start = 0
        index_in_epoch = batch_size
        assert batch_size <= num_examples
    end = index_in_epoch
    return x_train[start:end], y_train[start:end]  
    
    



# Create some wrappers for simplicity
def conv2d(x, W, b, strides=1):
    # Conv2D wrapper, with bias and relu activation
    x = tf.nn.conv2d(x, W, strides=[1, strides, strides, 1], padding='SAME')
    x = tf.nn.bias_add(x, b)
    return tf.nn.relu(x)


def maxpool2d(x, k=2):
    # MaxPool2D wrapper
    return tf.nn.max_pool(x, ksize=[1, k, k, 1], strides=[1, k, k, 1],padding='SAME')


# Create model
def model(x, weights, biases, dropout):
    # Reshape input picture
    x = tf.reshape(x, shape=[-1, 5, 8, 1])
    # Convolution Layer
    conv1 = conv2d(x, weights['wc1'], biases['bc1'])
    # Max Pooling (down-sampling)
    conv1 = maxpool2d(conv1, k=2)
    # Convolution Layer
    conv2 = conv2d(conv1, weights['wc2'], biases['bc2'])
    # Max Pooling (down-sampling)
    conv2 = maxpool2d(conv2, k=2)
    # Fully connected layer
    # Reshape conv2 output to fit fully connected layer input
    fc1 = tf.reshape(conv2, [-1, weights['wd1'].get_shape().as_list()[0]])
    fc1 = tf.add(tf.matmul(fc1, weights['wd1']), biases['bd1'])
    fc1 = tf.nn.relu(fc1)
    # Apply Dropout
    fc1 = tf.nn.dropout(fc1, dropout)
    # Output, class prediction
    out = tf.add(tf.matmul(fc1, weights['out']), biases['out'])
    return out

# Store layers weight & bias
weights = {
    # 5x5 conv, 1 input, 32 outputs
    'wc1': tf.Variable(tf.random_normal([5, 5, 1, 32])),
    # 5x5 conv, 32 inputs, 64 outputs
    'wc2': tf.Variable(tf.random_normal([5, 5, 32, 64])),
    # fully connected, 7*7*64 inputs, 1024 outputs
    'wd1': tf.Variable(tf.random_normal([7*7*64, 1024])),
    # 1024 inputs, 10 outputs (class prediction)
    'out': tf.Variable(tf.random_normal([1024, n_classes]))
}

biases = {
    'bc1': tf.Variable(tf.random_normal([32])),
    'bc2': tf.Variable(tf.random_normal([64])),
    'bd1': tf.Variable(tf.random_normal([1024])),
    'out': tf.Variable(tf.random_normal([n_classes]))
}


# Construct model
pred = model(x, weights, biases, keep_prob)

# Define loss and optimizer
cost = tf.reduce_mean(tf.nn.softmax_cross_entropy_with_logits(logits=pred, labels=y))
optimizer = tf.train.AdamOptimizer(learning_rate=learning_rate).minimize(cost)

# Evaluate model
correct_pred = tf.equal(tf.argmax(pred, 1), tf.argmax(y, 1))
accuracy = tf.reduce_mean(tf.cast(correct_pred, tf.float32))

# Initializing the variables
init = tf.global_variables_initializer()

# Launch the graph
with tf.Session() as sess:
    sess.run(init)
    step = 1
    # Keep training until reach max iterations
    while step * batch_size < training_iters:
        batchx, batchy = next_batch(batch_size)
        # Run optimization op (backprop)
        sess.run(optimizer, feed_dict={x: batchx , y: batchy,keep_prob: dropout})
        if step % display_step == 0:
            # Calculate batch loss and accuracy
            loss, acc = sess.run([cost, accuracy], feed_dict={x: batch_x, y: batch_y, keep_prob: 1.})
            print("Iter " + str(step*batch_size) + ", Minibatch Loss= " + \
                  "{:.6f}".format(loss) + ", Training Accuracy= " + \
                  "{:.5f}".format(acc))
            step += 1
        print("Optimization Finished!")

		
     # Calculate accuracy for 256 mnist test images
	print("Testing Accuracy:", \
    sess.run(accuracy, feed_dict={x: x_test, y: y_test, keep_prob: 1.}))
