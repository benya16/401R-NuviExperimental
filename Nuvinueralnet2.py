import tensorflow as tf
import random as rn
import pandas as pd
from sklearn.cross_validation import train_test_split
import numpy as np     

  
# Parameters
learning_rate = 0.001
training_iters = 200
batch_size = 1
display_step = 10



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

x_train,x_test,y_train,y_test = train_test_split(MNIST,train_labels,test_size=.1)
# we now have features (x_train) and y values, separated into test and train

# convert to dtype float 32
x_train,x_test,y_train,y_test = np.array(x_train,dtype='float32'), np.array(x_test,dtype='float32'),np.array(y_train,dtype='float32'),np.array(y_test,dtype='float32')


epochs_completed = 0
index_in_epoch = 0
num_examples = x_train.shape[0]
num_columns = x_train.shape[1]
hidden_units = 40
model_path = "C:/Users/Greg/Anaconda3/envs/tensorflow/Scripts/model.ckpt"
   

# Network Parameters
n_input = batch_size #  dta input (img shape: 1*10)
n_classes = 2 # total classes (0-1 digits)
dropout = 0.75 # Dropout, probability to keep units
hidden = .5

# tf Graph input
x = tf.placeholder(tf.float32, [1,num_columns])
y = tf.placeholder(tf.float32, [1, n_classes])
keep_prob = tf.placeholder(tf.float32) #dropout (keep probability)
keep_hidden = tf.placeholder(tf.float32)



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
    
	
def init_weights(shape):
    return tf.Variable(tf.random_normal(shape, stddev=0.01))


# Create model
def model(x,  w_h, w_h2, w_o, keep_prob,keep_hidden):
    x = tf.nn.dropout(x, keep_prob)
    h = tf.nn.relu(tf.matmul(x, w_h))
    h = tf.nn.dropout(h, keep_hidden)
    h2 = tf.nn.relu(tf.matmul(h, w_h2))
    h2 = tf.nn.dropout(h2, keep_hidden)
    return tf.matmul(h2, w_o)

	
	
# Store layers weight & bias
w_h = init_weights([num_columns, hidden_units])
w_h2 = init_weights([hidden_units,n_classes])
w_o = init_weights([n_classes, n_classes])

# Construct model
pred = model(x, w_h, w_h2, w_o, keep_prob, keep_hidden)

# Define loss and optimizer
cost = tf.reduce_mean(tf.nn.softmax_cross_entropy_with_logits(logits=pred, labels=y))
optimizer = tf.train.AdamOptimizer(learning_rate=learning_rate).minimize(cost)

# Evaluate model
correct_pred = tf.equal(tf.argmax(pred, 1), tf.argmax(y, 1))
accuracy = tf.reduce_mean(tf.cast(correct_pred, tf.float32))

# Initializing the variables
init = tf.global_variables_initializer()


saver = tf.train.Saver()
export_dir = "C:/Users/Greg/Anaconda3/envs/tensorflow/Scripts/model.ckpt"


# Launch the graph
with tf.Session() as sess:
    sess.run(init)
    step = 1
    step2 = 1
    correct = 0
    perc = 0
    # Keep training until reach max iterations
    while step < training_iters:
        batchx, batchy = next_batch(batch_size)
        # Run optimization op (backprop)
        sess.run(optimizer, feed_dict={x: batchx , y: batchy, keep_prob: dropout, keep_hidden: hidden})
        if step % display_step == 0:
            # Calculate batch loss and accuracy
            correct = 0
            perc = 0
            loss, acc = sess.run([cost, accuracy], feed_dict={x: batchx, y: batchy, keep_prob: 1., keep_hidden: hidden})
            correct = correct + acc
            perc = correct / step2
            print("Iter " + str(step*batch_size) + ", Minibatch Loss= " + "{:.6f}".format(loss) + "Training Accuracy= " + "{:.5f}".format(perc))
            step2 = 0
        step += 1
        step2 += 1
    print("Optimization Finished!")
     # Calculate accuracy for test data
    save_path = saver.save(sess, model_path)
    print("Model saved in file: %s" % save_path)
	
   
    	
	
	
def accuracy(predictions, labels):
      pred_class = np.argmax(predictions, 1)
      true_class = np.argmax(labels, 1)
      print (str(100.0 * np.sum(pred_class == true_class) / predictions.shape[0])



	

